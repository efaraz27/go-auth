package repositories

import (
	"github.com/efaraz27/go-auth/core"
	"github.com/efaraz27/go-auth/models"
	"gorm.io/gorm"
)

// UserRepository is a struct that defines the user repository
type UserRepository struct {
	db *gorm.DB
}

// NewUserRepository is a function that returns a new user repository
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

// Create is a method that creates a new user
func (r *UserRepository) Create(user *models.User) (*models.User, *core.Exception, error) {

	if err := r.db.Create(user).Error; err != nil {
		return nil, nil, err
	}

	return user, nil, nil
}

// FindAll is a method that returns all users
func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User

	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

// FindByUUID is a method that returns a user by UUID
func (r *UserRepository) FindByUUID(uuid string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("uuid = ?", uuid).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// FindByEmail is a method that returns a user by email
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User

	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update is a method that updates a user
func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Delete is a method that deletes a user
func (r *UserRepository) Delete(id int) error {
	if err := r.db.Delete(&models.User{}, id).Error; err != nil {
		return err
	}

	return nil
}
