package repositories

import (
	"context"
	errWrap "user-services/common/error"
	errConstant "user-services/constants/error"
	"user-services/domain/dto"
	"user-services/domain/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}


type IUserRepository interface {
	Register(context.Context, *dto.RegisterRequest) (*models.User, error)
	Update(context.Context, *dto.UpdateUserRequest, string) (*models.User, error)
	FindByUsername(context.Context, string) (*models.User, error)
	FindByEmail(context.Context, string) (*models.User, error)
	FindByUUID(context.Context, string) (*models.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Register(ctx context.Context, req *dto.RegisterRequest) (*models.User, error) {
	user := models.User{
		UUID:        uuid.New(),
		Name:        req.Name,
		Username:    req.Username,
		Password:    req.Password,
		PhoneNumber: req.PhoneNumber,
		Email:       req.Email,
		RoleID:      req.RoleID,
	}

	err := r.db.WithContext(ctx).Create(&user).Error
	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}

	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, req *dto.UpdateUserRequest, uuid string) (*models.User, error) {
	user := models.User{
		Name:        req.Name,
		Email: 	 req.Email,
		Username:   req.Username,
		Password:   req.Password,
		PhoneNumber: req.PhoneNumber,
	}
	err := r.db.WithContext(ctx).Model(&models.User{}).Where("uuid = ?", uuid).Updates(&user).Error

	if err != nil {
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}
	return &user, nil
}


func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}
	return &user, nil
}

func (r *UserRepository) FindByUUID(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Role").Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}
	return &user, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Preload("Role").Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errConstant.ErrUserNotFound
		}
		return nil, errWrap.WrapError(errConstant.ErrSQLError)
	}
	return &user, nil
}

