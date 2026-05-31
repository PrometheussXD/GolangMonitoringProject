package handlers

import (
	"Project/internal/models"
	"Project/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTarget(c *gin.Context) {
	var input models.Target

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	target := repository.CreateTarget(input)

	c.JSON(http.StatusOK, target)
}

func GetTargets(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetTargets())
}
