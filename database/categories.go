package db

import (
	"23f1002440KP/inventory-system/models"
)

type CreateCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Create A Category
func (database *SQLDB) CreateCategory(newCategory *CreateCategory) error {

	var category models.Category
	category.Name = newCategory.Name
	category.Description = newCategory.Description

	createdCategory := database.Db.Create(&category)
	if createdCategory.Error != nil {
		return createdCategory.Error
	}

	return nil

}

// Get all Categories
func (database *SQLDB) GetAllCategories() ([]models.Category, error) {

	var categories []models.Category
	err := database.Db.Find(&categories)
	if err.Error != nil {
		return nil, err.Error
	}

	return categories, nil
}

// Get a Category details by ID
func (database *SQLDB) GetCategoryByID(id string) (*models.Category, error) {

	var category models.Category
	err := database.Db.Find(&category, id)
	if err.Error != nil {
		return nil, err.Error
	}

	return &category, nil
}

// Update a Category details by ID
func (database *SQLDB) UpdateCategoryByID(id string, updatedCategory *CreateCategory) error {

	var category models.Category
	err := database.Db.Find(&category, id)
	if err.Error != nil {
		return err.Error
	}

	category.Name = updatedCategory.Name
	category.Description = updatedCategory.Description

	updated := database.Db.Save(&category)
	if updated.Error != nil {
		return updated.Error
	}

	return nil

}

// Delete a Category by ID
func (database *SQLDB) DeleteCategoryByID(id string) error {

	var category models.Category
	err := database.Db.Find(&category, id)
	if err.Error != nil {
		return err.Error
	}

	deleted := database.Db.Delete(&category)
	if deleted.Error != nil {
		return deleted.Error
	}

	return nil

}
