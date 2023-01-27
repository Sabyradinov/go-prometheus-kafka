package collector

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/segmentio/kafka-go"
)

var (
	readerDialCnt      = "kafka_reader_dial_count"
	readerFetchCnt     = "kafka_reader_fetch_count"
	readerMessageCnt   = "kafka_reader_message_count"
	readerMessageBytes = "kafka_reader_message_bytes"
	readerRebalanceCnt = "kafka_reader_rebalance_count"
	readerTimeoutCnt   = "kafka_reader_timeout_count"
	readerErrorCnt     = "kafka_reader_error_count"

	readerDialSecAvg = "kafka_reader_dial_seconds_avg"
	readerDialSecMin = "kafka_reader_dial_seconds_min"
	readerDialSecMax = "kafka_reader_dial_seconds_max"

	readerReadSecAvg = "kafka_reader_read_seconds_avg"
	readerReadSecMin = "kafka_reader_read_seconds_min"
	readerReadSecMax = "kafka_reader_read_seconds_max"

	readerWaitSecAvg = "kafka_reader_wait_seconds_avg"
	readerWaitSecMin = "kafka_reader_wait_seconds_min"
	readerWaitSecMax = "kafka_reader_wait_seconds_max"

	readerFetchSizeAvg = "kafka_reader_fetch_size_avg"
	readerFetchSizeMin = "kafka_reader_fetch_size_min"
	readerFetchSizeMax = "kafka_reader_fetch_size_max"

	readerFetchBytesAvg = "kafka_reader_fetch_bytes_avg"
	readerFetchBytesMin = "kafka_reader_fetch_bytes_min"
	readerFetchBytesMax = "kafka_reader_fetch_bytes_max"

	readerReaderOffset        = "kafka_reader_offset"
	readerReaderLag           = "kafka_reader_lag"
	readerReaderFetchBytesMin = "kafka_reader_fetch_bytes_min"
	readerReaderFetchBytesMax = "kafka_reader_fetch_bytes_max"
	readerReaderFetchWaitMax  = "kafka_reader_fetch_wait_max"
	readerReaderQueueLength   = "kafka_reader_queue_length"
	readerReaderQueueCapacity = "kafka_reader_queue_capacity"
)

type Reader interface {
	Stats() kafka.ReaderStats
}

type ReaderCollector struct {
	reader             Reader
	readerDialCnt      *prometheus.Desc
	readerFetchCnt     *prometheus.Desc
	readerMessageCnt   *prometheus.Desc
	readerMessageBytes *prometheus.Desc
	readerRebalanceCnt *prometheus.Desc
	readerTimeoutCnt   *prometheus.Desc
	readerErrorCnt     *prometheus.Desc

	readerDialSecAvg    *prometheus.Desc
	readerDialSecMin    *prometheus.Desc
	readerDialSecMax    *prometheus.Desc
	readerReadSecAvg    *prometheus.Desc
	readerReadSecMin    *prometheus.Desc
	readerReadSecMax    *prometheus.Desc
	readerWaitSecAvg    *prometheus.Desc
	readerWaitSecMin    *prometheus.Desc
	readerWaitSecMax    *prometheus.Desc
	readerFetchSizeAvg  *prometheus.Desc
	readerFetchSizeMin  *prometheus.Desc
	readerFetchSizeMax  *prometheus.Desc
	readerFetchBytesAvg *prometheus.Desc
	readerFetchBytesMin *prometheus.Desc
	readerFetchBytesMax *prometheus.Desc

	readerReaderOffset        *prometheus.Desc
	readerReaderLag           *prometheus.Desc
	readerReaderFetchBytesMin *prometheus.Desc
	readerReaderFetchBytesMax *prometheus.Desc
	readerReaderFetchWaitMax  *prometheus.Desc
	readerReaderQueueLength   *prometheus.Desc
	readerReaderQueueCapacity *prometheus.Desc
}

func NewReaderCollector(r Reader) *ReaderCollector {
	return &ReaderCollector{
		reader: r,
		//TODO check correct help in rus
		readerDialCnt: prometheus.NewDesc(
			readerDialCnt,
			"Dials count",
			nil,
			nil,
		),
		readerFetchCnt: prometheus.NewDesc(
			readerFetchCnt,
			"Fetched message count",
			nil,
			nil,
		),
		readerMessageCnt: prometheus.NewDesc(
			readerMessageCnt,
			"Received message count",
			nil,
			nil,
		),
		readerMessageBytes: prometheus.NewDesc(
			readerMessageBytes,
			"Read message byte size",
			nil,
			nil,
		),
		readerRebalanceCnt: prometheus.NewDesc(
			readerRebalanceCnt,
			"Rebalanced message count",
			nil,
			nil,
		),
		readerTimeoutCnt: prometheus.NewDesc(
			readerTimeoutCnt,
			"Timeouts count",
			nil,
			nil,
		),
		readerErrorCnt: prometheus.NewDesc(
			readerErrorCnt,
			"Errors count",
			nil,
			nil,
		),

		readerDialSecAvg: prometheus.NewDesc(
			readerDialSecAvg,
			"Average time for dial",
			nil,
			nil,
		),
		readerDialSecMin: prometheus.NewDesc(
			readerDialSecMin,
			"Min time for dial",
			nil,
			nil,
		),
		readerDialSecMax: prometheus.NewDesc(
			readerDialSecMax,
			"Max time for dial",
			nil,
			nil,
		),
		readerReadSecAvg: prometheus.NewDesc(
			readerReadSecAvg,
			"Average read time",
			nil,
			nil,
		),
		readerReadSecMin: prometheus.NewDesc(
			readerReadSecMin,
			"Min read time",
			nil,
			nil,
		),
		readerReadSecMax: prometheus.NewDesc(
			readerReadSecMax,
			"Max read time",
			nil,
			nil,
		),
		readerWaitSecAvg: prometheus.NewDesc(
			readerWaitSecAvg,
			"Average time for wait",
			nil,
			nil,
		),
		readerWaitSecMin: prometheus.NewDesc(
			readerWaitSecMin,
			"Min time for wait",
			nil,
			nil,
		),
		readerWaitSecMax: prometheus.NewDesc(
			readerWaitSecMax,
			"Max time for wait",
			nil,
			nil,
		),
		readerFetchSizeAvg: prometheus.NewDesc(
			readerFetchSizeAvg,
			"Average fetched size",
			nil,
			nil,
		),
		readerFetchSizeMin: prometheus.NewDesc(
			readerFetchSizeMin,
			"Min fetched size",
			nil,
			nil,
		),
		readerFetchSizeMax: prometheus.NewDesc(
			readerFetchSizeMax,
			"Max fetched size",
			nil,
			nil,
		),
		readerFetchBytesAvg: prometheus.NewDesc(
			readerFetchBytesAvg,
			"Average fetched bytes size",
			nil,
			nil,
		),
		readerFetchBytesMin: prometheus.NewDesc(
			readerFetchBytesMin,
			"Min fetched bytes size",
			nil,
			nil,
		),
		readerFetchBytesMax: prometheus.NewDesc(
			readerFetchBytesMax,
			"Max fetched bytes size",
			nil,
			nil,
		),
		readerReaderOffset: prometheus.NewDesc(
			readerReaderOffset,
			"Reader Offset",
			nil,
			nil,
		),
		readerReaderLag: prometheus.NewDesc(
			readerReaderLag,
			"Reader lags",
			nil,
			nil,
		),
		readerReaderFetchBytesMin: prometheus.NewDesc(
			readerReaderFetchBytesMin,
			"Min bytes size for fetch",
			nil,
			nil,
		),
		readerReaderFetchBytesMax: prometheus.NewDesc(
			readerReaderFetchBytesMax,
			"Max bytes size for fetch",
			nil,
			nil,
		),
		readerReaderFetchWaitMax: prometheus.NewDesc(
			readerReaderFetchWaitMax,
			"Max time for fetch",
			nil,
			nil,
		),
		readerReaderQueueLength: prometheus.NewDesc(
			readerReaderQueueLength,
			"Reader queue length",
			nil,
			nil,
		),
		readerReaderQueueCapacity: prometheus.NewDesc(
			readerReaderQueueCapacity,
			"Reader queue capacity",
			nil,
			nil,
		),
	}
}

func (rc ReaderCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- rc.readerDialCnt
	ch <- rc.readerFetchCnt
	ch <- rc.readerMessageCnt
	ch <- rc.readerMessageBytes
	ch <- rc.readerRebalanceCnt
	ch <- rc.readerTimeoutCnt
	ch <- rc.readerErrorCnt

	ch <- rc.readerDialSecAvg
	ch <- rc.readerDialSecMin
	ch <- rc.readerDialSecMax
	ch <- rc.readerReadSecAvg
	ch <- rc.readerReadSecMin
	ch <- rc.readerReadSecMax
	ch <- rc.readerWaitSecAvg
	ch <- rc.readerWaitSecMin
	ch <- rc.readerWaitSecMax
	ch <- rc.readerFetchSizeAvg
	ch <- rc.readerFetchSizeMin
	ch <- rc.readerFetchSizeMax
	ch <- rc.readerFetchBytesAvg
	ch <- rc.readerFetchBytesMin
	ch <- rc.readerFetchBytesMax
}

func (rc ReaderCollector) Collect(ch chan<- prometheus.Metric) {
	stats := rc.reader.Stats()

	ch <- prometheus.MustNewConstMetric(
		rc.readerDialCnt,
		prometheus.CounterValue,
		float64(stats.Dials),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchCnt,
		prometheus.CounterValue,
		float64(stats.Fetches),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerMessageCnt,
		prometheus.CounterValue,
		float64(stats.Messages),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerMessageBytes,
		prometheus.CounterValue,
		float64(stats.Bytes),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerRebalanceCnt,
		prometheus.CounterValue,
		float64(stats.Rebalances),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerTimeoutCnt,
		prometheus.CounterValue,
		float64(stats.Timeouts),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerErrorCnt,
		prometheus.CounterValue,
		float64(stats.Errors),
	)

	ch <- prometheus.MustNewConstMetric(
		rc.readerDialSecAvg,
		prometheus.GaugeValue,
		float64(stats.DialTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerDialSecMin,
		prometheus.GaugeValue,
		float64(stats.DialTime.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerDialSecMax,
		prometheus.GaugeValue,
		float64(stats.DialTime.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReadSecAvg,
		prometheus.GaugeValue,
		float64(stats.ReadTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReadSecMin,
		prometheus.GaugeValue,
		float64(stats.ReadTime.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReadSecMax,
		prometheus.GaugeValue,
		float64(stats.ReadTime.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerWaitSecAvg,
		prometheus.GaugeValue,
		float64(stats.WaitTime.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerWaitSecMin,
		prometheus.GaugeValue,
		float64(stats.WaitTime.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerWaitSecMax,
		prometheus.GaugeValue,
		float64(stats.WaitTime.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchSizeAvg,
		prometheus.GaugeValue,
		float64(stats.FetchSize.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchSizeMin,
		prometheus.GaugeValue,
		float64(stats.FetchSize.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchSizeMax,
		prometheus.GaugeValue,
		float64(stats.FetchSize.Max),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchBytesAvg,
		prometheus.GaugeValue,
		float64(stats.FetchBytes.Avg),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchBytesMin,
		prometheus.GaugeValue,
		float64(stats.FetchBytes.Min),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerFetchBytesMax,
		prometheus.GaugeValue,
		float64(stats.FetchBytes.Max),
	)

	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderOffset,
		prometheus.GaugeValue,
		float64(stats.Offset),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderLag,
		prometheus.GaugeValue,
		float64(stats.Lag),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderFetchBytesMin,
		prometheus.GaugeValue,
		float64(stats.MinBytes),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderFetchBytesMax,
		prometheus.GaugeValue,
		float64(stats.MaxBytes),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderFetchWaitMax,
		prometheus.GaugeValue,
		float64(stats.MaxWait),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderQueueLength,
		prometheus.GaugeValue,
		float64(stats.QueueLength),
	)
	ch <- prometheus.MustNewConstMetric(
		rc.readerReaderQueueCapacity,
		prometheus.GaugeValue,
		float64(stats.QueueCapacity),
	)
}
