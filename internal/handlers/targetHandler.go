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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := repository.CreateTarget(input)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, input)
}
func GetTargets(c *gin.Context) {

	targets, err := repository.GetTargets()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, targets)
}
