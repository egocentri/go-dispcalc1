package handlers

import (
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    "github.com/egocentri/go-dispcalc1/internal/models"
    "github.com/egocentri/go-dispcalc1/internal/orchestrator/services"
)

func GetTask(exprManager *services.ExpressionManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        task, ok := exprManager.GetNextTask()
        if !ok {
            c.JSON(http.StatusNotFound, gin.H{"error": "no tasks available"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"task": task})
    }
}

func PostTaskResult(exprManager *services.ExpressionManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        var result models.TaskResultRequest
        if err := c.ShouldBindJSON(&result); err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request"})
            return
        }

        if result.ID <= 0 {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid task id"})
            return
        }

        err := exprManager.SetTaskResult(strconv.Itoa(result.ID), result.Result)
        if err != nil {
            if err.Error() == "task not found" {
                c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
            } else {
                c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
            }
            return
        }

        c.JSON(http.StatusOK, gin.H{"status": "result accepted"})
    }
}
