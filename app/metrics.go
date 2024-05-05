package main

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	info     *prometheus.GaugeVec
	routines prometheus.Gauge
	latency  *prometheus.HistogramVec
	summary  *prometheus.SummaryVec
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		routines: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "connected_routines",
			Help:      "Number of currently connected routines",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "routine",
			Help:      "Information about the My App environment.",
		}, []string{"code", "id", "name"}),
		latency: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "app",
			Name:      "routine_latency",
			Help:      "Duration of the request.",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"id", "name", "path"}),
		summary: prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace: "app",
			Name:      "summary",
			Help:      "Duration of the login request.",
			Objectives: map[float64]float64{
				0.99: 0.001,
				0.9:  0.01,
				0.5:  0.05,
			},
		}, []string{"name"}),
	}

	reg.MustRegister(m.info, m.routines, m.latency, m.summary)
	return m
}
