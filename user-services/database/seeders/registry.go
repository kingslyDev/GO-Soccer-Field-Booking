package seeders

import "gorm.io/gorm"

type Registy struct {
	db *gorm.DB
}

type ISeederRegistry interface {
	Run()
}


func NewSeederRegistry(db *gorm.DB) ISeederRegistry {
	return &Registy{db: db}
} 


func (s *Registy) Run() {
	RunRoleSeeder(s.db)
	RunUserSeeder(s.db)
}