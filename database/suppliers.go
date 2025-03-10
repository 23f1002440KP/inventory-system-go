package db

import (
	"23f1002440KP/inventory-system/models"
)

//supplier Struct (Model)
type SupplierStruct struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone_number"`
	Address string `json:"address"`
}


//add a supplier (POST)
func (database *SQLDB) CreateSupplier(newSupplier *SupplierStruct) error {

	var supplier models.Supplier

	supplier.Name = newSupplier.Name
	supplier.Email = newSupplier.Email
	supplier.Phone = newSupplier.Phone
	supplier.Address = newSupplier.Address

	createdSupplier := database.Db.Create(&supplier)
	
	if createdSupplier.Error != nil {
		return createdSupplier.Error
	}

	return nil

}

//get all suppliers (GET)
func (database *SQLDB) GetAllSuppliers() ([]models.Supplier, error){
	var suppliers []models.Supplier
	err := database.Db.Find(&suppliers)
	if err.Error != nil {
		return nil, err.Error
	}
	return suppliers, nil
}

//get a supplier by id (GET)
func (database *SQLDB) GetSupplierByID(id string) (*models.Supplier, error){
	var supplier models.Supplier
	err := database.Db.Where("id=?",id).First(&supplier)
	if err.Error != nil {
		return nil, err.Error
	}
	return &supplier, nil
}

//update a supplier details by id (PUT)
func (database *SQLDB) UpdateSupplierByID(id string, updatedSupplier *SupplierStruct) error {
	var supplier models.Supplier

	err := database.Db.Where("id=?",id).First(&supplier)
	
	if err.Error != nil {
		return err.Error
	}

	supplier.Name = updatedSupplier.Name
	supplier.Email = updatedSupplier.Email
	supplier.Phone = updatedSupplier.Phone
	supplier.Address = updatedSupplier.Address

	updated := database.Db.Save(&supplier)
	if updated.Error != nil {
		return updated.Error
	}
	return nil

}

//delete a supplier by id (DELETE)

func (database *SQLDB) DeleteSupplierByID(id string) error {

	var supplier models.Supplier	
	err := database.Db.Where("id=?",id).First(&supplier)
	if err.Error != nil {
		return err.Error
	}

	supplier.IsDeleted = true
	// soft delete 
	deleted := database.Db.Save(&supplier)
	if deleted.Error != nil {
		return deleted.Error
	}

	return nil 
}
