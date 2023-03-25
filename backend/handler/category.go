package handler

import (
	"fmt"
	"net/http"
	"proyecto-integrador/dao"
	"proyecto-integrador/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddCategory handles POST requests and adds a new car product category
func AddCategory(c *gin.Context) {
	// parse database from request to category struct, bind JSON
	category := new(models.Category)
	if err := c.BindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// insert in db
	if _, err := dao.AddCategory(category); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// success response
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "success",
		"database": category,
	})
}

// ListAll handles GET requests and returns all car product categories
func ListAll(c *gin.Context) {
	categories := &[]models.Category{}
	// query db
	if err := dao.ListAll(categories); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "success",
		"database": categories,
	})
}

// GetCategoryByID handles GET requests and returns a car product category by id
func GetCategoryByID(c *gin.Context) {
	category := new(models.Category)
	// parse path param id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// query db
	isFound, err := dao.GetCategoryByID(id, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	// if !isFound, could not find category by id
	if !isFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Record with id %v not found.", id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "success",
		"database": category,
	})
}

// UpdateCategoryByID handles PATCH requests updates a car product category by id
func UpdateCategoryByID(c *gin.Context) {
	category := new(models.Category)
	// parse database from request body
	if err := c.BindJSON(category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// parse path param id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// update in db
	affected, err := dao.UpdateCategoryByID(id, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if affected == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Record with id %v not found.", id),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"msg":      "success",
		"database": category,
	})
}

// DeleteCategoryByID handles DELETE requests updates a car product category by id
func DeleteCategoryByID(c *gin.Context) {
	category := new(models.Category)
	// parse path param id
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	// query database
	affected, err := dao.DeleteCategoryByID(id, category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	if affected == -1 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Record with id %v not found.", id),
		})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
