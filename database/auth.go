package db

import (
	// "23f1002440KP/inventory-system/logger"
	"23f1002440KP/inventory-system/models"
	"23f1002440KP/inventory-system/utils"
	"fmt"
	// "golang.org/x/crypto/bcrypt"
)




func (database *SQLDB) Register(name string, email string, password string, role string) (int64, error) {

	if role != "admin" && role != "staff" {
		return 0, fmt.Errorf("role is not valid")
	}
	var roleUser models.Role

	if role == "admin" {
		roleUser = models.Admin
	} else if role == "staff" {
		roleUser = models.Staff
	}
	if validate := utils.ValidatePassword(password); !validate  {
		return 0, fmt.Errorf("password is not valid")
	}

	hashedPassword := utils.HashPasword(password)

	user := models.User{
		Name:          name,
		Email:         email,
		Password_hash: hashedPassword,
		Role:          roleUser,
	}
	if database == nil {
		return 0, fmt.Errorf("database is nil auth_interface")
	}
	newUser := database.Db.Create(&user)
	if newUser.Error != nil {
		return 0, newUser.Error
	}
	return int64(user.ID), nil

}


func (database *SQLDB) Login(email string, password string) (map[string]string, error) {
	var user models.User
	userPayload := make(map[string]string)

	if database == nil {
		return nil , fmt.Errorf("database is nil auth_interface")
	}

	database.Db.Where("email = ?", email).First(&user)


	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}
	hashedPassword := user.Password_hash
	binHash := []byte(hashedPassword)

	if !utils.CompareHashAndPassword(binHash, []byte(password)) {
		return nil , fmt.Errorf("password is not valid")
	}

	userPayload["id"] = fmt.Sprintf("%d", user.ID)
	userPayload["email"] = user.Email
	userPayload["role"] = string(user.Role)
	
	return userPayload , nil
	
}


func (database *SQLDB) GetUser(id string) (map[string]string, error) {
	var user models.User
	userPayload := make(map[string]string)

	if database == nil {
		return nil , fmt.Errorf("database is nil auth_interface")
	}
	database.Db.Where("id = ?", id).First(&user)

	if user.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	userPayload["name"] = user.Name
	userPayload["email"] = user.Email
	userPayload["role"] = string(user.Role)

	return userPayload , nil

}