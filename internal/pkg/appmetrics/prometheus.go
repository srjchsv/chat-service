package appmetrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type metricsCounters struct {
	HttpRequestsTotal  *prometheus.CounterVec
	WsChatMessagesSent prometheus.Counter
}

func InitPrometheus(r *gin.Engine) metricsCounters {
	metricsCounters := metricsCounters{}

	// Create a new Prometheus registry
	reg := prometheus.NewRegistry()
	metricsCounters.HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)
	metricsCounters.WsChatMessagesSent = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "websocket_messages_sent",
		Help: "Total number of WebSocket messages sent in chat",
	})
	// Create new prometheus registry
	reg.MustRegister(metricsCounters.HttpRequestsTotal, metricsCounters.WsChatMessagesSent)
	// Create a middleware that updates the HTTP requests counter
	r.Use(func(ctx *gin.Context) {
		// Call the next handler
		ctx.Next()
		// Update the HTTP requests counter
		metricsCounters.HttpRequestsTotal.WithLabelValues(
			ctx.Request.Method,
			ctx.Request.URL.Path,
			http.StatusText(ctx.Writer.Status()),
		).Inc()
	})
	// Register Prometheus metrics endpoint
	r.GET("/metrics", gin.WrapH(promhttp.HandlerFor(reg, promhttp.HandlerOpts{})))

	return metricsCounters
}

func IncCounter(conn prometheus.Counter) {
	conn.Inc()
}
