package dao

import (
	"proyecto-integrador/data"
	"proyecto-integrador/models"
)

// AddCategory inserts a new car product category in db
func AddCategory(c *models.Categories) (int64, error) {
	affected, err := data.DB.Insert(c)
	if err != nil {
		return -1, err
	}
	return affected, nil
}
