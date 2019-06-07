package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		status := c.Writer.Status()
		log.WithFields(log.Fields{
			"latency": latency.String(),
			"status":  status,
		}).Infoln()
	}
}
