package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/juliotorresmoreno/electric/handlers"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	handlers.AttachStatusHandler(r.Group("/status"))
	handlers.AttachConsumptionHandler(r.Group("/consumption"))

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8081
	r.Run(":8081")
}
