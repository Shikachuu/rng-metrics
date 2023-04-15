package pkg

import "github.com/prometheus/client_golang/prometheus"

type metrics struct {
	DebugCounter prometheus.Counter
	InfoCounter  prometheus.Counter
	WarnCounter  prometheus.Counter
	ErrorCounter prometheus.Counter
}

func NewMetrics() *metrics {
	return &metrics{
		DebugCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "log",
			Name:      "debug",
		}),
		InfoCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "log",
			Name:      "info",
		}),
		WarnCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "log",
			Name:      "warn",
		}),
		ErrorCounter: prometheus.NewCounter(prometheus.CounterOpts{
			Namespace: "log",
			Name:      "error",
		}),
	}
}
