package withSegmentio

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
	collector "go-prometheus-kafka"
)

func main() {
	w := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092", "localhost:9093", "localhost:9094"),
		Topic:    "topic-A",
		Balancer: &kafka.LeastBytes{},
	}
	writerCollector := collector.NewWriterCollector(w)
	prometheus.MustRegister(writerCollector)
}
