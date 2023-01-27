# go-prometheus-kafka
Prometheus metric exporter for kafka, only work with github.com/segmentio/kafka-go



Example with Consumer
```
package main

import (
	collector "github.com/Sabyradinov/go-prometheus-kafka"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
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
```

Example with Producer
```
package main

import (
	collector "github.com/Sabyradinov/go-prometheus-kafka"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
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

```