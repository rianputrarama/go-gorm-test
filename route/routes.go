package route

import (
	"gosqllearn/services"
	"gosqllearn/utils/logging"
	"os"
	"time"

	ginzap "github.com/gin-contrib/zap"

	"github.com/gin-gonic/gin"
)

func Init() {
	r := gin.Default()
	// set release mode
	gin.SetMode(gin.ReleaseMode)

	// error handler
	r.HandleMethodNotAllowed = true
	// r.NoMethod(service.MethodNotAllowed)
	// r.NoRoute(service.NotFound)

	start := time.Now()
	// set zap logger for gin router
	r.Use(ginzap.Ginzap(logging.HttpLogger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logging.HttpLogger, true))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/samples", services.GetDataSamples)
	}

	logging.Info("API service well initialized. Let's work!", "api service is ready", start)

	err := r.Run(os.Getenv("APP_PORT"))

	if err != nil {
		logging.Fatal(err.Error(), "router error", start)
	}
}
