package repository

import (
	"errors"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/lippyDesign/oath-service.git/entities"
)

// UserRepository interface to save user to DB
type UserRepository interface {
	FindOneByID(id int64) (*entities.User, error)
	FindOneByEmail(email string) (*entities.User, error)
	FindAll() ([]*entities.User, error)
	UpdateOne(user *entities.User) (*entities.User, error)
	CreateOne(user *entities.User) (*entities.User, error)
	DeleteOne(user *entities.User) error
}

type userRepo struct {
	db *gorm.DB
}

// newUserRepository initialize new user db client
func newUserRepository(db *gorm.DB) (UserRepository, error) {
	err := db.AutoMigrate(&entities.User{}).Error
	if err != nil {
		return nil, err
	}
	log.Println("Migrated User schema")
	var ur UserRepository = userRepo{db}
	return ur, nil
}

func (repo userRepo) FindAll() ([]*entities.User, error) {
	var users []*entities.User
	err := repo.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repo userRepo) FindOneByID(id int64) (*entities.User, error) {
	var user entities.User
	err := repo.db.Table("users").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo userRepo) FindOneByEmail(email string) (*entities.User, error) {
	var user entities.User
	err := repo.db.Table("users").Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo userRepo) CreateOne(user *entities.User) (*entities.User, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return nil, errors.New("Unable to create user with those credentials")
	}

	savedUser, err1 := repo.FindOneByEmail(user.Email)
	if err1 != nil {
		return nil, err1
	}
	return savedUser, nil
}

func (repo userRepo) UpdateOne(user *entities.User) (*entities.User, error) {
	err := repo.db.Save(&user).Error
	if err != nil {
		return nil, errors.New("Unable to update user with those credentials")
	}

	updatedUser, err1 := repo.FindOneByID(user.ID)
	if err1 != nil {
		return nil, err1
	}
	return updatedUser, nil
}

func (repo userRepo) DeleteOne(user *entities.User) error {
	if user.ID == 0 {
		err := errors.New("cannot delete user, missing id")
		return err
	}
	err := repo.db.Delete(&user).Error
	if err != nil {
		return err
	}
	return nil
}
