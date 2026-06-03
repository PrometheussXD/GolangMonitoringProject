package handlers

import (
	"Project/internal/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteTarget(c *gin.Context) {

	id, err :=
		strconv.Atoi(
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

	err =
		repository.DeleteTarget(
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

	c.Redirect(
		http.StatusFound,
		"/dashboard",
	)
}
