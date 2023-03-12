package dao

import (
	"proyecto-integrador/data"
	"proyecto-integrador/models"
)

// AddCategory inserts a new car product category
func AddCategory(c *models.Categories) (int64, error) {
	affected, err := data.DB.Insert(c)
	if err != nil {
		return -1, err
	}
	return affected, nil
}

// ListAll queries all car product categories
func ListAll(c *[]models.Categories) error {
	if err := data.DB.Find(c); err != nil {
		return err
	}
	return nil
}
