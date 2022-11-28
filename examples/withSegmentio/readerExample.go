package withSegmentio

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
	collector "go-prometheus-kafka"
)

func main() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092", "localhost:9093", "localhost:9094"},
		GroupID:  "consumer-group-id",
		Topic:    "topic-A",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	readerCollector := collector.NewReaderCollector(r)
	prometheus.MustRegister(readerCollector)
}
