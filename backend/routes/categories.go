package routes

import (
	"proyecto-integrador/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin router
	r := gin.Default()

	registerCategoriesRoutes(r)

	return r
}

var registerCategoriesRoutes = func(r *gin.Engine) {
	categories := r.Group("/categories")
	{
		categories.POST("/", handler.AddCategory)
		categories.GET("/", handler.ListAll)
	}
}
