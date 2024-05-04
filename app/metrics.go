package main

import "github.com/prometheus/client_golang/prometheus"

type Device struct {
	ID       int    `json:"id"`
	Mac      string `json:"mac"`
	Firmware string `json:"firmware"`
}

type metrics struct {
	devices       prometheus.Gauge
	info          *prometheus.GaugeVec
	upgrades      *prometheus.CounterVec
	duration      *prometheus.HistogramVec
	loginDuration prometheus.Summary
}

func NewMetrics(reg prometheus.Registerer) *metrics {
	m := &metrics{
		devices: prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "connected_devices",
			Help:      "Number of currently connected devices",
		}),
		info: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "app",
			Name:      "info",
			Help:      "Information about the My App environment.",
		}, []string{"version"}),
		upgrades: prometheus.NewCounterVec(prometheus.CounterOpts{
			Namespace: "app",
			Name:      "device_upgrade_total",
			Help:      "Number of upgraded devices.",
		}, []string{"type"}),
		duration: prometheus.NewHistogramVec(prometheus.HistogramOpts{
			Namespace: "app",
			Name:      "request_duration_seconds",
			Help:      "Duration of the request.",
			Buckets:   []float64{0.1, 0.15, 0.2, 0.25, 0.3},
		}, []string{"status", "method"}),
		loginDuration: prometheus.NewSummary(prometheus.SummaryOpts{
			Namespace:  "app",
			Name:       "login_request_duration_seconds",
			Help:       "Duration of the login request.",
			Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
		}),
	}

	reg.MustRegister(m.devices, m.info, m.upgrades, m.duration, m.loginDuration)
	return m
}
