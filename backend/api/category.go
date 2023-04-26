package api

import (
	"database/sql"
	"net/http"
	db "proyecto-integrador/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createCategoryRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
}

type getCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type listCategoryRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

type updateCategoryByIDRequestUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type updateCategoryByIDRequestBody struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageUrl    string `json:"image_url" binding:"required"`
}

type deleteCategoryRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) createCategory(c *gin.Context) {
	var req createCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateCategoryParams{
		Title:       req.Title,
		Description: req.Description,
		ImageUrl:    req.ImageUrl,
	}

	category, err := server.store.CreateCategory(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, category)
}

func (server *Server) getCategoryByID(c *gin.Context) {
	var req getCategoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	category, err := server.store.GetCategory(c, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, category)
}

func (server *Server) listCategory(c *gin.Context) {
	var req listCategoryRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListCategoriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	categories, err := server.store.ListCategories(c, arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (server *Server) updateCategoryByID(c *gin.Context) {
	var reqUri updateCategoryByIDRequestUri
	var reqBody updateCategoryByIDRequestBody
	if err := c.ShouldBindUri(&reqUri); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateCategoryParams{
		ID:          reqUri.ID,
		Title:       reqBody.Title,
		Description: reqBody.Description,
		ImageUrl:    reqBody.ImageUrl,
	}

	category, err := server.store.UpdateCategory(c, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, category)
}

func (server *Server) deleteCategoryByID(c *gin.Context) {
	var req deleteCategoryRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteCategory(c, req.ID)
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
