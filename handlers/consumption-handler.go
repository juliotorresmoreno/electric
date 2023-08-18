package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juliotorresmoreno/electric/db"
	"github.com/juliotorresmoreno/electric/models"
)

type ConsumptionHandler struct{}

type ConsumptionEntity struct {
	ID                 uint    `json:"id"`
	Address            string  `json:"address"`
	Period             string  `json:"period"`
	Active             float32 `json:"active"`
	ReactiveInductive  float32 `json:"reactive_inductive"`
	ReactiveCapacitive float32 `json:"reactive_capacitive"`
	Exported           float32 `json:"exported"`
}

type FindAllResponse struct {
	Period    []string    `json:"period"`
	DataGraph interface{} `json:"data_graph"`
}

func AttachConsumptionHandler(g *gin.RouterGroup) {
	h := &ConsumptionHandler{}
	g.GET("", h.FindAll)
}

func (h *ConsumptionHandler) FindAll(c *gin.Context) {
	data := make([]models.Consumption, 0)

	conn, err := db.GetConnection()
	if err != nil {
		c.JSON(500, ErrorResponse{Message: "Service is not working!"})
		return
	}
	tx := conn.Preload("Location").Find(&data)
	if tx.Error != nil {
		c.JSON(500, ErrorResponse{Message: "Service is not working!"})
		return
	}

	response := &FindAllResponse{
		Period:    []string{},
		DataGraph: []interface{}{},
	}

	c.JSON(http.StatusOK, response)
}

func (h *ConsumptionHandler) FindOne(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
