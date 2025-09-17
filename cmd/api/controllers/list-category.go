package controllers

import (
	"net/http"

	"github.com/alexandrewada/microservice-go-estudos/internal/repositories"
	"github.com/gin-gonic/gin"
)

func ListCategory(c *gin.Context, repository repositories.ICategoryRepository) {

	// Simulate category creation (in a real application, you would save to a database)
	categories, err := repository.List()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category listed successfully", "categories": categories})
}
