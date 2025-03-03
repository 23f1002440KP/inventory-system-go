package db

import (
	"23f1002440KP/inventory-system/logger"
	"23f1002440KP/inventory-system/models"
	"23f1002440KP/inventory-system/utils"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	logger.Logger().Info("Seeding the database")
	// Seed the database

	existingAdmin := models.User{
		Email: "Admin@inventSys.com",
	}
	
	existingAdminResult := db.Where("email = ?",existingAdmin.Email).Find(&existingAdmin)
	
	if existingAdminResult.RowsAffected > 0 {
		logger.Logger().Warn("Admin user already exists")
		} else {
			adminUser := models.User{
				Name:          "Admin",
				Email:         "Admin@inventSys.com",
				Password_hash: utils.HashPasword("Admin@123"),
				Role:          models.Admin,
		}
		logger.Logger().Warn("Creating the admin user")
		resultAdmin := db.Create(&adminUser)
		if resultAdmin.Error != nil {
			logger.Logger().Error("Not able to create the admin user")
			return
		}
		logger.Logger().Info("Admin user created")
		
	}

	existingStaff := models.User{
		Email: "Staff@inventSys.com",
	}
	existingStaffResult := db.Where("email = ?", existingStaff.Email).Find(&existingStaff)

	if existingStaffResult.RowsAffected > 0  {
		logger.Logger().Warn("Staff user already exists")
	} else {

		staffUser := models.User{
			Name:          "Staff",
			Email:         "Staff@inventSys.com",
			Password_hash: utils.HashPasword("Staff@123"),
			Role:          models.Staff,
		}

		logger.Logger().Warn("Creating the staff user")
		resultStaff := db.Create(&staffUser)
		if resultStaff.Error != nil {
			logger.Logger().Error("Not able to create the staff user")
			return
		}
		logger.Logger().Info("Staff user created")
	}
	logger.Logger().Info("Database seeded")

}
