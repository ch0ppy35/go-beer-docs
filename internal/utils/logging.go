package utils

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	l "github.com/sirupsen/logrus"
)

// SetupJsonLogger sets up the JSON log formatter
func SetupJsonLogger() {
	env := GetEnv("GIN_ENV", "dev")
	app := "svc_beer"

	l.SetFormatter(&JSONFormatter{
		App: app,
		Env: env,
	})
	l.SetLevel(l.DebugLevel)
}

// Timestamps in microsecond resolution
var timeStampFormat = "2006-01-02T15:04:05.000000Z07:00"

// JSONFormatter is a logger for use with Logrus
type JSONFormatter struct {
	App string
	Env string
}

func (f *JSONFormatter) Format(entry *l.Entry) ([]byte, error) {
	data := make(l.Fields, len(entry.Data)+3)
	for k, v := range entry.Data {
		data[k] = v
	}
	data["time"] = entry.Time.UTC().Format(timeStampFormat)
	data["msg"] = entry.Message
	data["level"] = strings.ToUpper(entry.Level.String())
	data["app"] = f.App
	data["env"] = f.Env

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}

// GetClientIP gets the correct IP for the end client instead of the proxy
func GetClientIP(c *gin.Context) string {
	requester := c.Request.Header.Get("X-Forwarded-For")
	if len(requester) == 0 {
		requester = c.Request.Header.Get("X-Real-IP")
	}
	if len(requester) == 0 {
		requester = c.Request.RemoteAddr
	}
	if strings.Contains(requester, ",") {
		requester = strings.Split(requester, ",")[0]
	}

	return requester
}

func GetDurationInMillseconds(start time.Time) float64 {
	end := time.Now()
	duration := end.Sub(start)
	milliseconds := float64(duration) / float64(time.Millisecond)
	rounded := float64(int(milliseconds*100+.5)) / 100
	return rounded
}
