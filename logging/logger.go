package logging

import (
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func SetupLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
		ForceColors:     true,
		DisableQuote:    true,
	})

	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetOutput(os.Stdout)
}

func RequestLogger() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		start := time.Now()
		reqID := uuid.New().String()[:8]

		c.Set("req_id", reqID)
		c.Header("X-Request-ID", reqID)

		logrus.WithFields(logrus.Fields{
			"req_id": reqID,
			"method": c.Request.Method,
			"path":   c.Request.URL.Path,
		}).Info("Start")
		c.Next()

		duration := time.Since(start)
		logrus.WithFields(logrus.Fields{
			"req_id":   reqID,
			"status":   c.Writer.Status(),
			"duration": duration,
		}).Info("End")
	})
}

func Log(c *gin.Context) *logrus.Entry {
	reqID, _ := c.Get("req_id")
	return logrus.WithField("req_id", reqID)
}
