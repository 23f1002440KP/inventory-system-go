package utils

import (
	"23f1002440KP/inventory-system/logger"

	"golang.org/x/crypto/bcrypt"
)

func HashPasword(password string) string {
	// hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		logger.Logger().Error("Not able to hash the password")
		return ""
	}

	return string(hashedPassword)
}

func CompareHashAndPassword(hashedPassword []byte, password []byte) bool {
	// compare the password
	var result bool = true
	// compare the password 
	// convert byte to string : string(hashedPassword)
	// convert string to byte : []byte(hashedPassword)
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		result = false		
	}

	return result
}