package repository

import (
	"log"

	"github.com/jinzhu/gorm"
)

// Repo object that has db methods
type Repo struct {
	User UserRepository
	DB   *gorm.DB
}

// NewRepo initializes DB
func NewRepo() (*Repo, error) {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return nil, err
	}
	r := Repo{}
	r.DB = db

	log.Println("Initalized DB")
	userRepository, errUserRepo := newUserRepository(db)
	if errUserRepo != nil {
		return nil, errUserRepo
	}
	r.User = userRepository
	return &r, nil
}
