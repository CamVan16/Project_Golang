package repositories

import (
	"camvan/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	SignUp(user *models.User) error
	SignIn(phone string) (models.User, error)
	//UpdateToken(user *models.User) error
	FindAll() ([]models.User, error)
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) SignUp(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) SignIn(phone string) (models.User, error) {
	var user models.User
	err := r.db.Where("Phone = ?", phone).First(&user).Error
	return user, err
}

func (r *userRepository) FindAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
