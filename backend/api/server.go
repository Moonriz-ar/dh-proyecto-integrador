package api

import (
	db "proyecto-integrador/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	Router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{
		store:  store,
		Router: gin.Default(),
	}

	// register routes to router
	server.registerRoutes()

	return server
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) registerRoutes() {
		server.Router.POST("/category", server.createCategory)
		server.Router.GET("/category", server.listCategory)
		server.Router.GET("/category/:id", server.getCategoryByID)
		server.Router.PATCH("/category/:id", server.updateCategoryByID)
		server.Router.DELETE("/category/:id", server.deleteCategoryByID)

		server.Router.POST("/city", server.createCity)
		server.Router.GET("/city", server.listCity)
		server.Router.GET("/city/:id", server.getCityByID)
	}