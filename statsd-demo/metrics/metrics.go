package metrics

import (
	"math/rand"
	"net/http"
	"regexp"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// prometheus 监控指标中间件

var (
	uptime        *prometheus.CounterVec   // 运行时间
	reqCount      *prometheus.CounterVec   // 请求数
	reqDuration   *prometheus.HistogramVec // 请求时间
	reqSizeBytes  *prometheus.SummaryVec   // 请求大小
	respSizeBytes *prometheus.SummaryVec   // 返回值大小
	roundGauge    *prometheus.GaugeVec     // 随机值
)

func randN() uint32 {
	return rand.Uint32() % 999
}

func RoundMetrics(status, endpoint, method string) {
	lvs := []string{status, endpoint, method}
	reqCount.WithLabelValues(lvs...).Inc()
	reqDuration.WithLabelValues(lvs...).Observe(float64(randN()))
	reqSizeBytes.WithLabelValues(lvs...).Observe(float64(randN()))
	respSizeBytes.WithLabelValues(lvs...).Observe(float64(randN()))
	roundGauge.WithLabelValues(lvs...).Set(float64(randN()))
}

// recordUptime increases service uptime per second.
func recordUptime() {
	for range time.Tick(time.Second) {
		uptime.WithLabelValues().Inc()
	}
}

// 设置指标
func registerMetrics(namespace string) {
	labels := []string{"status", "endpoint", "method"}
	uptime = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "uptime",
			Help:      "HTTP service uptime.",
		}, nil,
	)
	// 存活时间
	go recordUptime()
	reqCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: namespace,
			Name:      "http_request_count_total",
			Help:      "Total number of HTTP requests made.",
		}, labels,
	)

	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Name:      "http_request_duration_seconds",
			Help:      "HTTP request latencies in seconds.",
		}, labels,
	)

	reqSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_request_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)

	respSizeBytes = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Namespace: namespace,
			Name:      "http_response_size_bytes",
			Help:      "HTTP request sizes in bytes.",
		}, labels,
	)
	roundGauge = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace:   namespace,
			Name:        "round_gauge",
			Help:        "round gauge size",
			ConstLabels: nil,
		}, labels,
	)
	// 注册
	prometheus.MustRegister(uptime, reqCount, reqDuration, reqSizeBytes, respSizeBytes, roundGauge)
	return
}

// Init 初始化失败会抛出panic
func Init(namespace string) {
	registerMetrics(namespace)
}

// calcRequestSize returns the size of request object.
func calcRequestSize(r *http.Request) float64 {
	size := 0
	if r.URL != nil {
		size = len(r.URL.String())
	}

	size += len(r.Method)
	size += len(r.Proto)

	for name, values := range r.Header {
		size += len(name)
		for _, value := range values {
			size += len(value)
		}
	}
	size += len(r.Host)

	// r.Form and r.MultipartForm are assumed to be included in r.URL.
	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}
	return float64(size)
}

// PromOpts represents the Prometheus middleware Options.
// It was used for filtering labels with regex.
type PromOpts struct {
	ExcludeRegexStatus   string
	ExcludeRegexEndpoint string
	ExcludeRegexMethod   string
}

var defaultPromOpts = &PromOpts{}

// checkLabel returns the match result of labels.
// Return true if regex-pattern compiles failed.
func (po *PromOpts) checkLabel(label, pattern string) bool {
	if pattern == "" {
		return true
	}

	matched, err := regexp.MatchString(label, pattern)
	if err != nil {
		return true
	}
	return !matched
}
