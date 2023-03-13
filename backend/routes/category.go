package routes

import (
	"proyecto-integrador/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	// gin router
	r := gin.Default()

	registerCategoryRoutes(r)

	return r
}

var registerCategoryRoutes = func(r *gin.Engine) {
	category := r.Group("/category")
	{
		category.POST("/", handler.AddCategory)
		category.GET("/", handler.ListAll)
		category.GET("/:id", handler.GetCategoryByID)
		category.PATCH("/:id", handler.UpdateCategoryByID)
	}
}
