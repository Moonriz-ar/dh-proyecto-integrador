package dao

import (
	"proyecto-integrador/database"
	"proyecto-integrador/models"
)

// AddCategory inserts a new car product category
func AddCategory(c *models.Category) (int64, error) {
	affected, err := database.DB.Insert(c)
	if err != nil {
		return -1, err
	}
	return affected, nil
}

// ListAll queries all car product Category
func ListAll(c *[]models.Category) error {
	if err := database.DB.Find(c); err != nil {
		return err
	}
	return nil
}

// GetCategoryByID queries a car product category by id
func GetCategoryByID(id int, c *models.Category) (bool, error) {
	isFound, err := database.DB.ID(id).Get(c)
	if err != nil {
		return false, err
	}
	if !isFound {
		return isFound, nil
	}
	return isFound, nil
}

// UpdateCategoryByID updates a car product category by id
func UpdateCategoryByID(id int, c *models.Category) (int64, error) {
	affected, err := database.DB.ID(id).Update(c)
	if err != nil {
		return -1, err
	}
	// if affected == 0, means there is no category with that id
	if affected == 0 {
		return -1, nil
	}
	// retrieve updated record
	if _, err := database.DB.ID(id).Get(c); err != nil {
		return affected, err
	}
	return affected, nil
}

// DeleteCategoryByID deletes a car product category by id
func DeleteCategoryByID(id int, c *models.Category) (int64, error) {
	affected, err := database.DB.ID(id).Delete(c)
	if err != nil {
		return -1, err
	}
	// affected == 0 means there is no category with that id
	if affected == 0 {
		return -1, nil
	}
	// affected == 1 means category with that id has been deleted
	return affected, nil
}
