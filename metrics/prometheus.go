package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	RampUpTime = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "ramp_up_time",
			Help: "Time it took for service to be ready (ns).",
		},
	)

	hdFailures = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "hd_errors_total",
			Help: "Number of hard-disk errors.",
		},
		[]string{"device"},
	)
)

func init() {
	// Metrics have to be registered to be exposed:
	prometheus.MustRegister(RampUpTime)
	prometheus.MustRegister(hdFailures)

	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()
}
