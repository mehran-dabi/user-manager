package repository

import (
	"context"
	"database/sql"
	"faceit/domain/constants"
	"faceit/domain/user/entity"
	"faceit/domain/user/utils"
	"fmt"

	"github.com/go-redis/redis"
)

type IUsersRepository interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	Update(ctx context.Context, user *entity.User) error
	Remove(ctx context.Context, ID int64) error
	GetByID(ctx context.Context, ID int64) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByNickName(ctx context.Context, nickName string) (*entity.User, error)
	Get(ctx context.Context, filter *entity.Filter, page, pageSize int64) ([]*entity.User, error)
	GetCount(ctx context.Context, filter *entity.Filter) (uint64, error)
	PublishUserChangeEvent(userID int64) error
}

type UsersRepository struct {
	db    *sql.DB
	redis redis.UniversalClient
}

func NewUserRepository(db *sql.DB, redis redis.UniversalClient) *UsersRepository {
	return &UsersRepository{db: db, redis: redis}
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
	query := utils.UpdateQueryBuilder(user, usersTableName)
	_, err := u.db.ExecContext(
		ctx,
		query,
	)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if err := u.PublishUserChangeEvent(user.ID); err != nil {
		return err
	}

	return nil
}

// Remove - removes the user with the given ID
func (u *UsersRepository) Remove(ctx context.Context, ID int64) error {
	result, err := u.db.ExecContext(
		ctx,
		deleteUser,
		ID,
	)
	if err != nil {
		return fmt.Errorf("failed to remove user: %w", err)
	}

	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get number of rows affected: %w", err)
	}

	if count == 0 {
		return fmt.Errorf("no users were deleted")
	}

	return nil
}

// GetByID - gets the user from database with the given ID
func (u *UsersRepository) GetByID(ctx context.Context, ID int64) (*entity.User, error) {
	result, err := u.db.QueryContext(
		ctx,
		getUserByID,
		ID,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	defer func(result *sql.Rows) {
		_ = result.Close()
	}(result)

	if !result.Next() {
		return nil, constants.ErrUserNotFound
	}

	user := &entity.User{}
	if err := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Email,
		&user.Country,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to read user from database: %w", err)
	}

	return user, nil
}

// GetByEmail - gets the user from database with the given email
func (u *UsersRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	result, err := u.db.QueryContext(
		ctx,
		getUserByEmail,
		email,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	defer func(result *sql.Rows) {
		_ = result.Close()
	}(result)

	if !result.Next() {
		return nil, constants.ErrUserNotFound
	}

	user := &entity.User{}
	if err := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Email,
		&user.Country,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to read user from database: %w", err)
	}

	return user, nil
}

// GetByNickName - gets the user from database with the given nick name
func (u *UsersRepository) GetByNickName(ctx context.Context, nickName string) (*entity.User, error) {
	result, err := u.db.QueryContext(
		ctx,
		getUserByNickName,
		nickName,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to query database: %w", err)
	}

	defer func(result *sql.Rows) {
		_ = result.Close()
	}(result)

	if !result.Next() {
		return nil, constants.ErrUserNotFound
	}

	user := &entity.User{}
	if err := result.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.NickName,
		&user.Email,
		&user.Country,
		&user.CreatedAt,
		&user.UpdatedAt,
	); err != nil {
		return nil, fmt.Errorf("failed to read user from database: %w", err)
	}

	return user, nil
}

// Get - return the users with the provided criteria in the filter field and return the data with pagination and the total count of the results.
func (u *UsersRepository) Get(ctx context.Context, filter *entity.Filter, page, pageSize int64) ([]*entity.User, error) {
	query := utils.QueryBuilder(filter, usersTableName, page, pageSize)

	results, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}

	defer func(results *sql.Rows) {
		_ = results.Close()
	}(results)

	var users []*entity.User
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
		); err != nil {
			return nil, fmt.Errorf("failed to read records from database: %w", err)
		}

		users = append(users, user)
	}

	return users, nil
}

// GetCount - gets the total count of users with the provided filter
func (u *UsersRepository) GetCount(ctx context.Context, filter *entity.Filter) (uint64, error) {
	query := utils.CountQueryBuilder(filter, usersTableName)

	result, err := u.db.QueryContext(ctx, query)
	if err != nil {
		return 0, fmt.Errorf("failed to get total count of users: %w", err)
	}

	defer func(results *sql.Rows) {
		_ = results.Close()
	}(result)

	if !result.Next() {
		return 0, constants.ErrUserNotFound
	}

	var count uint64
	if err := result.Scan(&count); err != nil {
		return 0, fmt.Errorf("failed to read record from database: %w", err)
	}

	return count, nil
}

// PublishUserChangeEvent - This function is used to store the user change event in the redis so other services can be notified of the change.
func (u *UsersRepository) PublishUserChangeEvent(userID int64) error {
	_, err := u.redis.RPush(UserChangesRedisKey, userID).Result()
	if err != nil {
		return fmt.Errorf("failed to push event to redis: %w", err)
	}
	return nil
}
