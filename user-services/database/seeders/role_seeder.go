package seeders

import (
	"user-services/domain/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func RunRoleSeeder(db *gorm.DB) {
	// code here
	role := []models.Role{
		{
			Code: "ADMIN",
			Name: "Administator",
		},
		{
			Code: "CUSTOMER",
			Name: "Customer",
		},
	}

	for i, role := range role {
		err := db.FirstOrCreate(&role, models.Role{Code: role.Code}).Error
		if err != nil {
			logrus.Errorf("Error while seeding role %d: %v", i, err)
			panic(err)
		}
		logrus.Infof("Role %s successfully seeded", role.Code)
	}
}