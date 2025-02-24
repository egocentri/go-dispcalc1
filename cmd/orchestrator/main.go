package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/egocentri/go-dispcalc1/internal/orchestrator/handlers"
	"github.com/egocentri/go-dispcalc1/internal/orchestrator/services"

)

func main() {
	services.InitTaskManager()

	router := gin.Default()
	router.POST("/api/v1/calculate", handlers.CalculateExpression)
	router.GET("/api/v1/expressions", handlers.GetExpressions)
	router.GET("/api/v1/expressions/:id", handlers.GetExpressionByID)
	router.GET("/internal/task", handlers.GetTask)
	router.POST("/internal/task", handlers.PostTaskResult)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Orchestrator started on :%s", port)
	router.Run(":" + port)
}
