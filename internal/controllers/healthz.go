package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthzController struct{}

func NewHealthzController(e *gin.Engine) {
	h := HealthzController{}
	e.GET("/healthz", h.GetHealthz)
}

func (h *HealthzController) GetHealthz(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
