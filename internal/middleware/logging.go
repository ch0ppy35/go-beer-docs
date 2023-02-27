package middleware

import (
	"time"

	"github.com/ch0ppy35/beer-docs/internal/utils"
	"github.com/gin-gonic/gin"
	l "github.com/sirupsen/logrus"
)

// JSONLogMiddleware logs a gin HTTP request in JSON
func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := utils.GetDurationInMillseconds(start)
		entry := l.WithFields(l.Fields{
			"client_ip":  utils.GetClientIP(c),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     c.Writer.Status(),
			"referrer":   c.Request.Referer(),
			"request_id": c.Writer.Header().Get("Request-Id"),
		})

		if c.Writer.Status() >= 500 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}
