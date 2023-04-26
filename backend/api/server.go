package api

import (
	db "proyecto-integrador/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store  db.Store
	Router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		Store:  store,
		Router: gin.Default(),
	}

	// register routes to router
	registerCategoryRoutes(server)
	registerCityRoutes(server)

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func registerCategoryRoutes(server *Server) {
	category := server.Router.Group("/category")
	{
		category.POST("/", server.createCategory)
		category.GET("/", server.listCategory)
		category.GET("/:id", server.getCategoryByID)
		category.PATCH("/:id", server.updateCategoryByID)
		category.DELETE("/:id", server.deleteCategoryByID)
	}
}

func registerCityRoutes(server *Server) {
	city := server.Router.Group("/city")
	{
		city.POST("/", server.createCity)
		city.GET("/", server.listCity)
		city.GET(":id", server.getCityByID)
	}
}
