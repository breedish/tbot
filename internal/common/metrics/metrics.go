package metrics

import (
	"github.com/breedish/tbot/internal/common/service"
	"github.com/prometheus/client_golang/prometheus"
)

type appMetricsCollector struct {
	actions *prometheus.CounterVec
}

func (m appMetricsCollector) Inc(action string, count int) {
	m.actions.With(prometheus.Labels{"action": action}).Add(float64(count))
}

func NewApplicationMetricsCollector(serviceInfo service.Info, registry prometheus.Registerer) *appMetricsCollector {
	m := &appMetricsCollector{
		actions: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: serviceInfo.ServiceName,
			Name:      "actions_total",
			Help:      "Number of processed actions.",
		}, []string{"type"}),
	}

	registry.MustRegister(m.actions)
	return m
}
