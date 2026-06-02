package handlers

import (
	"Project/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TargetDetails(c *gin.Context) {

	id, err := strconv.Atoi(
		c.Param("id"),
	)

	if err != nil {

		c.JSON(
			http.StatusBadRequest,
			gin.H{
				"error": "invalid id",
			},
		)

		return
	}

	results, err :=
		repository.GetResultsByTargetID(
			uint(id),
		)

	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"error": err.Error(),
			},
		)

		return
	}

	c.HTML(
		http.StatusOK,
		"target_details.html",
		gin.H{
			"Results": results,
		},
	)

}
