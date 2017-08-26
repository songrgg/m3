// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package msgpack

import (
	"io"
	"net"
	"sync"

	"github.com/m3db/m3aggregator/aggregator"
	"github.com/m3db/m3metrics/protocol/msgpack"
	"github.com/m3db/m3x/log"
	"github.com/m3db/m3x/server"

	"github.com/uber-go/tally"
)

const (
	unknownRemoteHostAddress = "<unknown>"
)

// NewServer creates a new msgpack server.
func NewServer(address string, aggregator aggregator.Aggregator, opts Options) xserver.Server {
	iOpts := opts.InstrumentOptions()
	handlerScope := iOpts.MetricsScope().Tagged(map[string]string{"handler": "msgpack"})
	handler := NewHandler(aggregator, opts.SetInstrumentOptions(iOpts.SetMetricsScope(handlerScope)))
	return xserver.NewServer(address, handler, opts.ServerOptions())
}

type handlerMetrics struct {
	addMetricErrors tally.Counter
	decodeErrors    tally.Counter
}

func newHandlerMetrics(scope tally.Scope) handlerMetrics {
	return handlerMetrics{
		addMetricErrors: scope.Counter("add-metric-errors"),
		decodeErrors:    scope.Counter("decode-errors"),
	}
}

type handler struct {
	sync.Mutex

	aggregator   aggregator.Aggregator
	opts         Options
	log          xlog.Logger
	iteratorPool msgpack.UnaggregatedIteratorPool
	metrics      handlerMetrics
}

// NewHandler creates a new msgpack handler.
func NewHandler(aggregator aggregator.Aggregator, opts Options) xserver.Handler {
	iOpts := opts.InstrumentOptions()
	return &handler{
		aggregator:   aggregator,
		opts:         opts,
		log:          iOpts.Logger(),
		iteratorPool: opts.IteratorPool(),
		metrics:      newHandlerMetrics(iOpts.MetricsScope()),
	}
}

func (s *handler) Handle(conn net.Conn) {
	it := s.iteratorPool.Get()
	it.Reset(conn)
	defer it.Close()

	// Iterate over the incoming metrics stream and queue up metrics.
	for it.Next() {
		metric := it.Metric()
		policiesList := it.PoliciesList()
		if err := s.aggregator.AddMetricWithPoliciesList(metric, policiesList); err != nil {
			s.log.WithFields(
				xlog.NewLogField("metric", metric.String()),
				xlog.NewLogField("policies", policiesList),
				xlog.NewLogErrField(err),
			).Errorf("error adding metric with policies")
			s.metrics.addMetricErrors.Inc(1)
		}
	}

	// If there is an error during decoding, it's likely due to a broken connection
	// and therefore we ignore the EOF error.
	if err := it.Err(); err != nil && err != io.EOF {
		remoteAddress := unknownRemoteHostAddress
		if remoteAddr := conn.RemoteAddr(); remoteAddr != nil {
			remoteAddress = remoteAddr.String()
		}
		s.log.WithFields(
			xlog.NewLogField("remoteAddress", remoteAddress),
			xlog.NewLogErrField(err),
		).Error("decode error")
		s.metrics.decodeErrors.Inc(1)
	}
}

func (s *handler) Close() {
	// NB(cw) Do not close s.aggregator here because it's shared between
	// the msgpack server and the http server, and it will be closed on
	// exit signal.
}