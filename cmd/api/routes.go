package main

import (
	"github.com/alexandrewada/microservice-go-estudos/cmd/api/controllers"
	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(router *gin.Engine) {
	categoryGroup := router.Group("/categories")

	categoryGroup.POST("/", controllers.CreateCategory)
}
