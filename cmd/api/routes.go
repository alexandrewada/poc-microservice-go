package main

import (
	"github.com/alexandrewada/microservice-go-estudos/cmd/api/controllers"
	"github.com/alexandrewada/microservice-go-estudos/internal/repositories"
	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/categories")

	inMemoryRepo := repositories.NewInMemoryCategoryRepository()

	categoryGroup.POST("/", func(c *gin.Context) {
		controllers.CreateCategory(c, inMemoryRepo)
	})

	categoryGroup.GET("/", func(c *gin.Context) {
		controllers.ListCategory(c, inMemoryRepo)
	})
}
