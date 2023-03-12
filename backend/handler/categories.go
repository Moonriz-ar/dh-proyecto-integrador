package handler

import (
	"net/http"
	"proyecto-integrador/dao"
	"proyecto-integrador/models"

	"github.com/gin-gonic/gin"
)

// AddCategory adds a new car product category
func AddCategory(c *gin.Context) {
	// parse data from request to category struct, bind JSON
	category := new(models.Categories)
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
		"code": 200,
		"msg":  "success",
		"data": category,
	})
}
