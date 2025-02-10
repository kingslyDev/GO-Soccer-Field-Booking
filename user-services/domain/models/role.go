package models

import "time"

type Role struct {
	ID uint `gorm:"primaryKey;autoIncrement"`
	Code string `gorm:"type:varchar(15);not null"`
	Name string `gorm:"type:varchar(50);not null"`
	CreateAt time.Time 
	UpdateAt time.Time


}