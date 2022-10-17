package repository

import "database/sql"

type IUsersRepository interface {
	Create()
}

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}
