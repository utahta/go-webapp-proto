package model

import "github.com/utahta/go-webapp-proto/app/lib/db"

// e.g.
//   repo := model.NewUserRepository()
//   user, err := repo.Find(1)
//

type userRepository struct{}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (repo *userRepository) Find(id int) (*User, error) {
	var user User
	if _, err := db.E().Id(id).Get(&user); err != nil {
		return nil, err
	}
	return &user, nil
}
