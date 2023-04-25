package api

import (
	db "proyecto-integrador/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		Store:  store,
		router: gin.Default(),
	}

	// register routes to router
	registerCategoryRoutes(server)

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func registerCategoryRoutes(server *Server) {
	category := server.router.Group("/category")
	{
		category.POST("/", server.createCategory)
		category.GET("/", server.listCategory)
		category.GET("/:id", server.getCategoryByID)
		category.PATCH("/:id", server.updateCategoryByID)
		category.DELETE("/:id", server.deleteCategoryByID)
	}
}
