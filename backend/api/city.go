package api

import (
	"database/sql"
	"fmt"
	"net/http"
	db "proyecto-integrador/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createCityRequest struct {
	Name string `json:"name" binding:"required"`
}

type getCityRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listCityRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) createCity(c *gin.Context) {
	var req createCityRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	city, err := server.store.CreateCity(c, req.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, city)
}

func (server *Server) getCityByID(c *gin.Context) {
	var req getCityRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	city, err := server.store.GetCity(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, city)
}

func (server *Server) listCity(c *gin.Context) {
	var req listCityRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		fmt.Println(req)
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCitiesParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	cities, err := server.store.ListCities(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, cities)
}