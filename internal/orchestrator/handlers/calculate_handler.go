package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/egocentri/go-dispcalc1/internal/models"
    "github.com/egocentri/go-dispcalc1/internal/orchestrator/services"
)

func CalculateExpression(exprManager *services.ExpressionManager) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req models.CalculateRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "invalid request body"})
            return
        }

        exprID, err := exprManager.CreateExpression(req.Expression)
        if err != nil {
            c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
            return
        }
        c.JSON(http.StatusCreated, gin.H{"id": exprID})
    }
}
