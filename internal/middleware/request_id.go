package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type RequestIDOptions struct {
	AllowSetting bool
}

func RequestID(options RequestIDOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string
		if options.AllowSetting {
			requestID = c.Request.Header.Get("Set-Request-Id")
		}
		if requestID == "" {
			requestID = uuid.NewString()
		}

		c.Writer.Header().Set("Request-Id", requestID)
		c.Next()
	}
}
