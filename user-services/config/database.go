package config

import (
	"fmt"
	"net/url"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDatabase() (*gorm.DB, error) {
	config := config
	endcodedPassword := url.QueryEscape(config.Database.Password)
	uri := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", 
		config.Database.Username, 
		endcodedPassword, 
		config.Database.Host, 
		config.Database.Port, 
		config.Database.Name,
	)
	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(config.Database.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(config.Database.MaxOpenConnection)
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database.MaxLifetimeConnection) * time.Second)
	sqlDB.SetConnMaxIdleTime(time.Duration(config.Database.MaxIdleTime) * time.Second)

	return db, nil
}