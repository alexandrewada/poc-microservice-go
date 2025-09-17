package controllers

import (
	"net/http"

	"github.com/alexandrewada/microservice-go-estudos/internal/entities"
	"github.com/alexandrewada/microservice-go-estudos/internal/repositories"
	use_cases "github.com/alexandrewada/microservice-go-estudos/internal/use-cases"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(c *gin.Context, repository repositories.ICategoryRepository) {
	var body createCategoryInput

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	// Simulate category creation (in a real application, you would save to a database)
	category, err := entities.NewCategory(body.Name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "fail",
			"error":  err.Error(),
		})
		return
	}

	useCase := use_cases.NewcreateCategoryUseCase(repository)
	err = useCase.Execute(body.Name)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status": "error",
			"error":  "Could not create category",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully", "category": category})
}
