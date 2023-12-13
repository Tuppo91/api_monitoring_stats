package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	MetricServiceTimeHistogramVec = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tonstatus_functions_time",
			Help:    "Service functions execution duration distribution in seconds",
			Buckets: []float64{0.01, 0.05, 0.1, 0.2, 0.3, 0.4, 0.5, 0.7, 1, 2.5, 5, 10},
		},
		[]string{"service"},
	)

	MetricServiceIndexingLatencyHistogramVec = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "tonstatus_indexing_latency",
			Help:    "difference between current time and last transaction on electror",
			Buckets: []float64{0.5, 1, 5, 7.5, 10, 12.5, 15, 17.5, 20, 25, 30, 60, 120, 300, 600, 1200, 3600},
		},
		[]string{"service"},
	)

	MetricServiceRequest = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "tonstatus_http_api",
		Help: "availability of http api",
	}, []string{"service"})
)
