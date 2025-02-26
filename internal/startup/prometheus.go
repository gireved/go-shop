package startup

import (
	"fmt"
	"go-shop/config"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// InitPrometheus 初始化 Prometheus 监控
func InitPrometheus() {
	prometheusConfig := config.Config.Prometheus

	// 创建默认 Prometheus 注册器
	registry := prometheus.NewRegistry()

	// 注册 Go 运行时和进程监控指标
	registry.MustRegister(
		collectors.NewGoCollector(),                                       // 监控 Go 运行时信息（替换废弃方法）
		collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}), // 监控进程信息
	)

	// 启动 HTTP 服务器
	go func() {
		http.Handle(prometheusConfig.Path, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

		address := fmt.Sprintf("%s:%d", prometheusConfig.Host, prometheusConfig.Port)
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalf("Prometheus 监控启动失败: %v", err)
		}

		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalf("Prometheus 监控启动失败: %v", err)
		}
	}()
}
