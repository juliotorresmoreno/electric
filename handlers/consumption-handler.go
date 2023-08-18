package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ConsumptionHandler struct{}

type FindAllResponse struct {
	Period    []string
	DataGraph []interface{}
}

func AttachConsumptionHandler(g *gin.RouterGroup) {
	h := &ConsumptionHandler{}
	g.GET("", h.FindAll)
	g.GET("/:id", h.FindAll)
}

func (h *ConsumptionHandler) FindAll(c *gin.Context) {
	response := new(FindAllResponse)
	c.JSON(http.StatusOK, response)
}

func (h *ConsumptionHandler) FindOne(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
