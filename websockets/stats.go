package websockets

import (
	"errors"

	"go.k6.io/k6/js/modules"
	"go.k6.io/k6/metrics"
)

type wsMetrics struct {
	wsBytesSent          *metrics.Metric
	wsBytesReceived      *metrics.Metric
	wsCurrentConnections *metrics.Metric
	wsFailedHandshakes   *metrics.Metric
	wsForcedDisconnects  *metrics.Metric
}

func registerMetrics(vu modules.VU) (wsMetrics, error) {
	var err error
	var m wsMetrics
	registry := vu.InitEnv().Registry

	if m.wsBytesSent, err = registry.NewMetric("ws_bytes_sent", metrics.Counter, metrics.Data); err != nil {
		return m, errors.Unwrap(err)
	}

	if m.wsBytesReceived, err = registry.NewMetric("ws_bytes_received", metrics.Counter, metrics.Data); err != nil {
		return m, errors.Unwrap(err)
	}

	if m.wsCurrentConnections, err = registry.NewMetric("ws_current_connections", metrics.Gauge, metrics.Data); err != nil {
		return m, errors.Unwrap(err)
	}

	if m.wsFailedHandshakes, err = registry.NewMetric("ws_failed_handshakes", metrics.Counter); err != nil {
		return m, errors.Unwrap(err)
	}

	if m.wsForcedDisconnects, err = registry.NewMetric("ws_forced_disconnects", metrics.Counter); err != nil {
		return m, errors.Unwrap(err)
	}

	return m, nil
}
