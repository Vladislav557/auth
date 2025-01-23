package resources

import (
	"fmt"
	"github.com/Vladislav557/auth/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
	"time"
)

var noLogging = []string{
	"/health",
	"/metrics",
}

func RouterInit() *gin.Engine {
	r := gin.New()
	configureRouter(r)
	addMetrics(r)
	addHandlers(r)
	return r
}

func addHandlers(r *gin.Engine) {
	rc := handlers.RegistrationHandler{}
	r.POST("/registration", rc.Register)
}

func addMetrics(r *gin.Engine) {
	m := ginmetrics.GetMonitor()
	m.SetMetricPath("/metrics")
	m.Use(r)
}

func configureRouter(r *gin.Engine) {
	formatter := func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s | %s | %d | %s \"%s\" | %s | %s | %s\n",
			param.TimeStamp.Format(time.RFC3339),
			param.Latency,
			param.StatusCode,
			param.Method,
			param.Path,
			param.ClientIP,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}
	gin.SetMode(gin.ReleaseMode)
	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: noLogging,
		Formatter: formatter,
	}),
		gin.Recovery(),
	)
}
