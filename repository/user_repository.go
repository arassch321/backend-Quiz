package repository

import (
	"quiz/models"

	"gorm.io/gorm"
)

type RepositoryUser interface {
	//create User
	Save(user *models.User) (*models.User, error)
	FindById(ID int) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
	Delete(user *models.User) (*models.User, error)
}

type repository_user struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repository_user {
	return &repository_user{db}
}

func (r *repository_user) Save(user *models.User) (*models.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) FindByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.db.Where("email = ?", email).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) FindById(ID int) (*models.User, error) {
	var user *models.User

	err := r.db.Where("id = ?", ID).Find(&user).Error

	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository_user) Delete(user *models.User) (*models.User, error) {
	err := r.db.Delete(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}
