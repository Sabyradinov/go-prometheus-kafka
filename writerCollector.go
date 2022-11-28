package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

const (
	DefaultNamespace = "go_prometheues_kafka"
	DefaultSubsystem = "kafka_stats"
)

type Writer interface {
	Stats() kafka.WriterStats
}

type WriterCollector struct {
	writer    Writer
	dialTime  *prometheus.Desc
	messages  *prometheus.Desc
	retries   *prometheus.Desc
	writeTime *prometheus.Desc
	errors    *prometheus.Desc
}

func NewWriterCollector(w Writer) *WriterCollector {
	return &WriterCollector{
		writer: w,
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
		retries: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "retries"),
			"Number of retries to connect server",
			nil,
			nil,
		),
		writeTime: prometheus.NewDesc(
			prometheus.BuildFQName(DefaultNamespace, DefaultSubsystem, "writeTime"),
			"Avg time of write messages to kafka server",
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

func (wc WriterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- wc.dialTime
	ch <- wc.messages
	ch <- wc.retries
	ch <- wc.writeTime
	ch <- wc.errors
}

func (wc WriterCollector) Collect(ch chan<- prometheus.Metric) {
	stats := wc.writer.Stats()

	ch <- prometheus.MustNewConstMetric(
		wc.dialTime,
		prometheus.CounterValue,
		float64(stats.DialTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.messages,
		prometheus.CounterValue,
		float64(stats.Messages),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.retries,
		prometheus.CounterValue,
		float64(stats.Retries),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writeTime,
		prometheus.CounterValue,
		float64(stats.WriteTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.errors,
		prometheus.CounterValue,
		float64(stats.Errors),
	)
}
