package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsumptionHandler struct{}

func AttachConsumptionHandler(g *gin.RouterGroup) {
	h := &ConsumptionHandler{}
	g.GET("", h.Get)
}

func (h *ConsumptionHandler) Get(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
