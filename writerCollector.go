package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

var (
	writerWriteCnt     = "kafka_writer_write_count"
	writerMessageCnt   = "kafka_writer_message_count"
	writerMessageBytes = "kafka_writer_message_bytes"
	writerErrorCnt     = "kafka_writer_error_count"

	writerBatchSecAvg = "kafka_writer_batch_seconds_avg"
	writerBatchSecMin = "kafka_writer_batch_seconds_min"
	writerBatchSecMax = "kafka_writer_batch_seconds_max"

	writerWriteSecAvg = "kafka_writer_write_seconds_avg"
	writerWriteSecMin = "kafka_writer_write_seconds_min"
	writerWriteSecMax = "kafka_writer_write_seconds_max"

	writerWaitSecAvg = "kafka_writer_wait_seconds_avg"
	writerWaitSecMin = "kafka_writer_wait_seconds_min"
	writerWaitSecMax = "kafka_writer_wait_seconds_max"

	writerRetriesCnt = "kafka_writer_retries_count"

	writerBatchSizeAvg = "kafka_writer_batch_size_avg"
	writerBatchSizeMin = "kafka_writer_batch_size_min"
	writerBatchSizeMax = "kafka_writer_batch_size_max"

	writerBatchBytesAvg = "kafka_writer_batch_bytes_avg"
	writerBatchBytesMin = "kafka_writer_batch_bytes_min"
	writerBatchBytesMax = "kafka_writer_batch_bytes_max"

	writerAttemptsMax  = "kafka_writer_attempts_max"
	writerBackoffMin   = "kafka_writer_backoff_min"
	writerBackoffMax   = "kafka_writer_backoff_max"
	writerBatchMax     = "kafka_writer_batch_max"
	writerBatchTimeout = "kafka_writer_batch_timeout"
	writerReadTimeout  = "kafka_writer_read_timeout"
	writerWriteTimeout = "kafka_writer_write_timeout"
	writerAcksRequired = "kafka_writer_acks_required"
)

type Writer interface {
	Stats() kafka.WriterStats
}

type WriterCollector struct {
	writer              Writer
	writerWriteCnt      *prometheus.Desc
	writerMessageCnt    *prometheus.Desc
	writerMessageBytes  *prometheus.Desc
	writerErrorCnt      *prometheus.Desc
	writerBatchSecAvg   *prometheus.Desc
	writerBatchSecMin   *prometheus.Desc
	writerBatchSecMax   *prometheus.Desc
	writerWriteSecAvg   *prometheus.Desc
	writerWriteSecMin   *prometheus.Desc
	writerWriteSecMax   *prometheus.Desc
	writerWaitSecAvg    *prometheus.Desc
	writerWaitSecMin    *prometheus.Desc
	writerWaitSecMax    *prometheus.Desc
	writerRetriesCnt    *prometheus.Desc
	writerBatchSizeAvg  *prometheus.Desc
	writerBatchSizeMin  *prometheus.Desc
	writerBatchSizeMax  *prometheus.Desc
	writerBatchBytesAvg *prometheus.Desc
	writerBatchBytesMin *prometheus.Desc
	writerBatchBytesMax *prometheus.Desc
	writerAttemptsMax   *prometheus.Desc
	writerBackoffMin    *prometheus.Desc
	writerBackoffMax    *prometheus.Desc
	writerBatchMax      *prometheus.Desc
	writerBatchTimeout  *prometheus.Desc
	writerReadTimeout   *prometheus.Desc
	writerWriteTimeout  *prometheus.Desc
	writerAcksRequired  *prometheus.Desc
}

func NewWriterCollector(w Writer) *WriterCollector {
	return &WriterCollector{
		writer: w,
		writerWriteCnt: prometheus.NewDesc(
			writerWriteCnt,
			"Produced messages count",
			nil,
			nil,
		),
		writerMessageCnt: prometheus.NewDesc(
			writerMessageCnt,
			"writerMessageCnt",
			nil,
			nil,
		),
		writerMessageBytes: prometheus.NewDesc(
			writerMessageBytes,
			"writerMessageBytes",
			nil,
			nil,
		),
		writerErrorCnt: prometheus.NewDesc(
			writerErrorCnt,
			"writerErrorCnt",
			nil,
			nil,
		),
		writerBatchSecAvg: prometheus.NewDesc(
			writerBatchSecAvg,
			"writerBatchSecAvg",
			nil,
			nil,
		),
		writerBatchSecMin: prometheus.NewDesc(
			writerBatchSecMin,
			"writerBatchSecMin",
			nil,
			nil,
		),
		writerBatchSecMax: prometheus.NewDesc(
			writerBatchSecMax,
			"writerBatchSecMax",
			nil,
			nil,
		),
		writerWriteSecAvg: prometheus.NewDesc(
			writerWriteSecAvg,
			"writerWriteSecAvg",
			nil,
			nil,
		),
		writerWriteSecMin: prometheus.NewDesc(
			writerWriteSecMin,
			"writerWriteSecMin",
			nil,
			nil,
		),
		writerWriteSecMax: prometheus.NewDesc(
			writerWriteSecMax,
			"writerWriteSecMax",
			nil,
			nil,
		),
		writerWaitSecAvg: prometheus.NewDesc(
			writerWaitSecAvg,
			"writerWaitSecAvg",
			nil,
			nil,
		),
		writerWaitSecMin: prometheus.NewDesc(
			writerWaitSecMin,
			"writerWaitSecMin",
			nil,
			nil,
		),
		writerWaitSecMax: prometheus.NewDesc(
			writerWaitSecMax,
			"writerWaitSecMax",
			nil,
			nil,
		),
		writerRetriesCnt: prometheus.NewDesc(
			writerRetriesCnt,
			"writerRetriesCnt",
			nil,
			nil,
		),
		writerBatchSizeAvg: prometheus.NewDesc(
			writerBatchSizeAvg,
			"writerBatchSizeAvg",
			nil,
			nil,
		),
		writerBatchSizeMin: prometheus.NewDesc(
			writerBatchSizeMin,
			"writerBatchSizeMin",
			nil,
			nil,
		),
		writerBatchSizeMax: prometheus.NewDesc(
			writerBatchSizeMax,
			"writerBatchSizeMax",
			nil,
			nil,
		),
		writerBatchBytesAvg: prometheus.NewDesc(
			writerBatchBytesAvg,
			"writerBatchBytesAvg",
			nil,
			nil,
		),
		writerBatchBytesMin: prometheus.NewDesc(
			writerBatchBytesMin,
			"writerBatchBytesMin",
			nil,
			nil,
		),
		writerBatchBytesMax: prometheus.NewDesc(
			writerBatchBytesMax,
			"writerBatchBytesMax",
			nil,
			nil,
		),
		writerAttemptsMax: prometheus.NewDesc(
			writerAttemptsMax,
			"writerAttemptsMax",
			nil,
			nil,
		),
		writerBackoffMin: prometheus.NewDesc(
			writerBackoffMin,
			"writerBackoffMin",
			nil,
			nil,
		),
		writerBackoffMax: prometheus.NewDesc(
			writerBackoffMax,
			"writerBackoffMax",
			nil,
			nil,
		),
		writerBatchMax: prometheus.NewDesc(
			writerBatchMax,
			"writerBatchMax",
			nil,
			nil,
		),
		writerBatchTimeout: prometheus.NewDesc(
			writerBatchTimeout,
			"writerBatchTimeout",
			nil,
			nil,
		),
		writerReadTimeout: prometheus.NewDesc(
			writerReadTimeout,
			"writerReadTimeout",
			nil,
			nil,
		),
		writerWriteTimeout: prometheus.NewDesc(
			writerWriteTimeout,
			"writerWriteTimeout",
			nil,
			nil,
		),
		writerAcksRequired: prometheus.NewDesc(
			writerAcksRequired,
			"writerAcksRequired",
			nil,
			nil,
		),
	}
}

func (wc WriterCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- wc.writerWriteCnt
	ch <- wc.writerMessageCnt
	ch <- wc.writerMessageBytes
	ch <- wc.writerErrorCnt
	ch <- wc.writerBatchSecAvg
	ch <- wc.writerBatchSecMin
	ch <- wc.writerBatchSecMax
	ch <- wc.writerWriteSecAvg
	ch <- wc.writerWriteSecMin
	ch <- wc.writerWriteSecMax
	ch <- wc.writerWaitSecAvg
	ch <- wc.writerWaitSecMin
	ch <- wc.writerWaitSecMax
	ch <- wc.writerRetriesCnt
	ch <- wc.writerBatchSizeAvg
	ch <- wc.writerBatchSizeMin
	ch <- wc.writerBatchSizeMax
	ch <- wc.writerBatchBytesAvg
	ch <- wc.writerBatchBytesMin
	ch <- wc.writerBatchBytesMax
	ch <- wc.writerAttemptsMax
	ch <- wc.writerBackoffMin
	ch <- wc.writerBackoffMax
	ch <- wc.writerBatchMax
	ch <- wc.writerBatchTimeout
	ch <- wc.writerReadTimeout
	ch <- wc.writerWriteTimeout
	ch <- wc.writerAcksRequired
}

func (wc WriterCollector) Collect(ch chan<- prometheus.Metric) {
	stats := wc.writer.Stats()

	ch <- prometheus.MustNewConstMetric(
		wc.writerWriteCnt,
		prometheus.CounterValue,
		float64(stats.Writes),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerMessageCnt,
		prometheus.CounterValue,
		float64(stats.Messages),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerMessageBytes,
		prometheus.CounterValue,
		float64(stats.Bytes),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerErrorCnt,
		prometheus.CounterValue,
		float64(stats.Errors),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchSecAvg,
		prometheus.GaugeValue,
		float64(stats.BatchTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchSecMin,
		prometheus.GaugeValue,
		float64(stats.BatchTime.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchSecMax,
		prometheus.GaugeValue,
		float64(stats.BatchTime.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWriteSecAvg,
		prometheus.GaugeValue,
		float64(stats.WriteTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWriteSecMin,
		prometheus.GaugeValue,
		float64(stats.WriteTime.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWriteSecMax,
		prometheus.GaugeValue,
		float64(stats.WriteTime.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWaitSecAvg,
		prometheus.GaugeValue,
		float64(stats.WaitTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWaitSecMin,
		prometheus.GaugeValue,
		float64(stats.WaitTime.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWaitSecMax,
		prometheus.GaugeValue,
		float64(stats.WaitTime.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerRetriesCnt,
		prometheus.CounterValue,
		float64(stats.Retries),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchSizeAvg,
		prometheus.GaugeValue,
		float64(stats.BatchSize.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchSizeMin,
		prometheus.GaugeValue,
		float64(stats.BatchSize.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchSizeMax,
		prometheus.GaugeValue,
		float64(stats.BatchSize.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchBytesAvg,
		prometheus.GaugeValue,
		float64(stats.BatchBytes.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchBytesMin,
		prometheus.GaugeValue,
		float64(stats.BatchBytes.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchBytesMax,
		prometheus.GaugeValue,
		float64(stats.BatchBytes.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerAttemptsMax,
		prometheus.GaugeValue,
		float64(stats.MaxAttempts),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBackoffMin,
		prometheus.GaugeValue,
		float64(stats.WriteBackoffMin),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBackoffMax,
		prometheus.GaugeValue,
		float64(stats.WriteBackoffMax),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchMax,
		prometheus.GaugeValue,
		float64(stats.MaxBatchSize),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerBatchTimeout,
		prometheus.GaugeValue,
		float64(stats.BatchTimeout),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerReadTimeout,
		prometheus.GaugeValue,
		float64(stats.ReadTimeout),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerWriteTimeout,
		prometheus.GaugeValue,
		float64(stats.WriteTimeout),
	)
	ch <- prometheus.MustNewConstMetric(
		wc.writerAcksRequired,
		prometheus.GaugeValue,
		float64(stats.RequiredAcks),
	)
}
