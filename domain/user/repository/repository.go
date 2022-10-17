package repository

import (
	"context"
	"database/sql"
	"faceit/domain/entity"
	"faceit/domain/utils"
	"fmt"
)

type IUsersRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Remove(ctx context.Context, ID int64) error
	Get(ctx context.Context, filter *entity.Filter, page, pageSize int64) ([]*entity.User, int64, error)
}

type UsersRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{db: db}
}

// Create - creates a user with the given information
func (u *UsersRepository) Create(ctx context.Context, user *entity.User) (*entity.User, error) {
	result, err := u.db.ExecContext(
		ctx,
		createUser,
		user.FirstName,
		user.LastName,
		user.NickName,
		user.Password,
		user.Email,
		user.Country,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	user.ID, err = result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed to get last inserted ID: %w", err)
	}

	return user, nil
}

// Update - updates the user with the given information
func (u *UsersRepository) Update(ctx context.Context, user *entity.User) error {
	_, err := u.db.ExecContext(
		ctx,
		updateUser,
		user.FirstName,
		user.LastName,
		user.NickName,
		user.Password,
		user.Email,
		user.Country,
		user.ID,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return err
}

// Remove - removes the user with the given ID
func (u *UsersRepository) Remove(ctx context.Context, ID int64) error {
	_, err := u.db.ExecContext(
		ctx,
		deleteUser,
		ID,
	)
	if err != nil {
		return fmt.Errorf("failed to remove user: %w", err)
	}

	return nil
}

// Get - return the users with the provided criteria in the filter field and return the data with pagination and the total count of the results.
func (u *UsersRepository) Get(ctx context.Context, filter *entity.Filter, page, pageSize int64) ([]*entity.User, int64, error) {
	query := utils.QueryBuilder(filter, usersTableName, page, pageSize)

	results, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get users: %w", err)
	}

	defer func(results *sql.Rows) {
		_ = results.Close()
	}(results)

	var users []*entity.User
	var count int64
	for results.Next() {
		user := new(entity.User)
		if err := results.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.NickName,
			&user.Email,
			&user.Country,
			&user.CreatedAt,
			&user.UpdatedAt,
			&count,
		); err != nil {
			return nil, 0, fmt.Errorf("failed to read records from database: %w", err)
		}

		users = append(users, user)
	}

	return users, count, nil
}
