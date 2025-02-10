package repositories

import (
	repositories "user-services/repositories/user"

	"gorm.io/gorm"
)



type Registry struct {
	db *gorm.DB
}

type IRepositoriesRegistry interface {
	GetUser() repositories.IUserRepository
}

func NewRepositoriesRegistry(db *gorm.DB) IRepositoriesRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetUser() repositories.IUserRepository {
	return repositories.NewUserRepository(r.db)
}

