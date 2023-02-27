package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CORSOptions struct {
	Origin string
}

// CORS middleware
func CORS(options CORSOptions) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // allow any origin domain
		if options.Origin != "" {
			c.Writer.Header().Set("Access-Control-Allow-Origin", options.Origin)
		}
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		} else {
			c.Next()
		}
	}
}
