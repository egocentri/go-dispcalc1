package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/egocentri/go-dispcalc1/internal/orchestrator/services"
)
func GetTask(exprManager *services.ExpressionManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		timeout := time.After(30 * time.Second)
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-timeout:
				c.JSON(http.StatusNotFound, gin.H{"error": "no tasks available"})
				return
			case <-ticker.C:
				task, ok := exprManager.GetNextTask()
				if ok {
					c.JSON(http.StatusOK, gin.H{"task": task})
					return
				}
			}
		}
	}
}
