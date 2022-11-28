package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

type Reader interface {
	Stats() kafka.ReaderStats
}

type ReaderCollector struct {
	reader   Reader
	dialTime *prometheus.Desc
	messages *prometheus.Desc
	lag      *prometheus.Desc
	readTime *prometheus.Desc
	errors   *prometheus.Desc
}

func NewReaderCollector(r Reader) *ReaderCollector {
	return &ReaderCollector{
		reader: r,
		dialTime: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "dialTime"),
			"Avg connection time",
			nil,
			nil,
		),
		messages: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "messages"),
			"Number of messages sent by writer",
			nil,
			nil,
		),
		lag: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "retries"),
			"Number of lags",
			nil,
			nil,
		),
		readTime: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "writeTime"),
			"Avg time to read messages from kafka server",
			nil,
			nil,
		),
		errors: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "errors"),
			"Number of error occurred in kafka writer",
			nil,
			nil,
		),
	}
}

func (rc ReaderCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- rc.dialTime
	ch <- rc.messages
	ch <- rc.lag
	ch <- rc.readTime
	ch <- rc.errors
}

func (rc ReaderCollector) Collect(ch chan<- prometheus.Metric) {
	stats := rc.reader.Stats()

	ch <- prometheus.MustNewConstMetric(
		rc.dialTime,
		prometheus.CounterValue,
		float64(stats.DialTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.messages,
		prometheus.CounterValue,
		float64(stats.Messages),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.lag,
		prometheus.CounterValue,
		float64(stats.Lag),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readTime,
		prometheus.CounterValue,
		float64(stats.ReadTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.errors,
		prometheus.CounterValue,
		float64(stats.Errors),
	)
}
