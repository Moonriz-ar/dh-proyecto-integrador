package api

import (
	"database/sql"
	"net/http"
	db "proyecto-integrador/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createProductRequest struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID int64 `json:"category_id" binding:"required"`
	CityID int64 `json:"city_id" binding:"required"`
}

type getProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listProductRequest struct {
	PageID int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type updateProductByIDRequestUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateProductByIDRequestBody struct {
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CategoryID int64 `json:"category_id" binding:"required"`
	CityID int64 `json:"city_id" binding:"required"`
}

type deleteProductRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createProduct(c *gin.Context) {
	var req createProductRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateProductParams{
		Title: req.Title,
		Description: req.Description,
		CategoryID: req.CategoryID,
		CityID: req.CityID,
	}

	product, err := server.store.CreateProduct(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, product)
}

func (server *Server) getProductByID(c *gin.Context) {
	var req getProductRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.store.GetProduct(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, product)
}

func (server *Server) listProduct(c *gin.Context) {
	var req listProductRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListProductParams{
		Limit: req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	products, err := server.store.ListProduct(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, products)
}

func (server *Server) updateProductByID(c *gin.Context) {
	var reqUri updateProductByIDRequestUri
	var reqBody updateProductByIDRequestBody
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateProductParams{
		ID: reqUri.ID,
		Title: reqBody.Title,
		Description: reqBody.Description,
		CategoryID: reqBody.CategoryID,
		CityID: reqBody.CityID,
	}

	product, err := server.store.UpdateProduct(c, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, product)
}

func (server *Server) deleteProductByID(c *gin.Context) {
	var req deleteProductRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteProduct(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusNoContent, nil)
}