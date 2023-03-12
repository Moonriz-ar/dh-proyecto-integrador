package dao

import (
	"proyecto-integrador/data"
	"proyecto-integrador/models"
)

// AddCategory inserts a new car product category
func AddCategory(c *models.Category) (int64, error) {
	affected, err := data.DB.Insert(c)
	if err != nil {
		return -1, err
	}
	return affected, nil
}

// ListAll queries all car product Category
func ListAll(c *[]models.Category) error {
	if err := data.DB.Find(c); err != nil {
		return err
	}
	return nil
}

// GetCategoryByID queries a car product category by id
func GetCategoryByID(id int, c *models.Category) (bool, error) {
	isFound, err := data.DB.ID(id).Get(c)
	if err != nil {
		return isFound, err
	}
	if !isFound {
		return isFound, err
	}
	return isFound, err
}
