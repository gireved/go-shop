package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"sync"
	"time"
)

type Builder struct {
	Namespace  string
	Subsystem  string
	Name       string
	InstanceId string
	Help       string
}

var (
	responseTimeVec    *prometheus.SummaryVec // 监控 HTTP 响应时间
	activeRequestGauge prometheus.Gauge       // 监控活跃请求数
	onceResponseTime   sync.Once
	onceActiveRequest  sync.Once
)

// BuildResponseTime 监控 HTTP 响应时间
func (b *Builder) BuildResponseTime() gin.HandlerFunc {
	onceResponseTime.Do(func() {
		// 以 {method, pattern, status} 作为 标签（Label），区分不同的 HTTP 方法、路由和状态码。
		labels := []string{"method", "pattern", "status"}
		responseTimeVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
			Namespace: b.Namespace,
			Subsystem: b.Subsystem,
			Name:      b.Name + "_resp_time",
			Help:      b.Help,
			ConstLabels: map[string]string{
				"instance_id": b.InstanceId,
			},
			Objectives: map[float64]float64{
				0.5:   0.01,
				0.75:  0.01,
				0.9:   0.01,
				0.99:  0.001,
				0.999: 0.0001,
			},
		}, labels)
		prometheus.MustRegister(responseTimeVec)
	})

	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		duration := time.Since(start).Seconds()

		// 获取请求方法、路径（pattern）和响应状态码
		method := ctx.Request.Method
		pattern := ctx.FullPath()
		if pattern == "" {
			pattern = ctx.Request.URL.Path // 兜底方案
		}
		status := strconv.Itoa(ctx.Writer.Status())

		// 记录请求的响应时间
		responseTimeVec.WithLabelValues(method, pattern, status).Observe(duration)
	}
}

// BuildActiveRequest 监控活跃请求数
func (b *Builder) BuildActiveRequest() gin.HandlerFunc {
	onceActiveRequest.Do(func() {
		activeRequestGauge = prometheus.NewGauge(prometheus.GaugeOpts{
			Namespace: b.Namespace,
			Subsystem: b.Subsystem,
			Name:      b.Name + "_active_req",
			Help:      b.Help,
			ConstLabels: map[string]string{
				"instance_id": b.InstanceId,
			},
		})
		prometheus.MustRegister(activeRequestGauge)
	})

	return func(ctx *gin.Context) {
		activeRequestGauge.Inc()
		defer activeRequestGauge.Dec()
		ctx.Next()
	}
}
