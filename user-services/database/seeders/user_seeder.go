package seeders

import (
	"user-services/constants"
	"user-services/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func RunUserSeeder(db *gorm.DB) {
	// code here
	password, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	user := models.User{
		UUID : uuid.New(),
		Name : "Administator",
		Username : "admin",
		Password: string(password),
		PhoneNumber: "08123456789",
		Email: "admin@gmail.com",
		RoleID: constants.Admin,
		
	}

	err := db.FirstOrCreate(&user, models.User{Username: user.Username}).Error
	if err != nil {
		logrus.Errorf("Error while seeding user %v", err)
		panic(err)
	}
	logrus.Infof("User %s successfully seeded", user.Username)
}