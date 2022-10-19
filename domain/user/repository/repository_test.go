package repository

import (
	"context"
	"database/sql"
	"faceit/domain/user/entity"
	databaseMocks "faceit/mocks/infrastructure/database"
	redisMocks "faceit/mocks/infrastructure/redis"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/alicebob/miniredis/v2"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RepositoryTestSuite struct {
	suite.Suite
	db    *sql.DB
	mock  sqlmock.Sqlmock
	redis *miniredis.Miniredis
}

func (r *RepositoryTestSuite) TestCreate() {
	testCases := []struct {
		user               *entity.User
		ctx                context.Context
		expectedUserEntity *entity.User
		expectedError      error
	}{
		{
			user: &entity.User{
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Password:  "pass",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			ctx: context.Background(),
			expectedUserEntity: &entity.User{
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Password:  "pass",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		r.mock.ExpectExec("INSERT INTO").
			WithArgs(tc.user.FirstName, tc.user.LastName, tc.user.NickName, tc.user.Password, tc.user.Email, tc.user.Country).
			WillReturnResult(sqlmock.NewResult(0, 1))
		userEntity, err := userRepository.Create(tc.ctx, tc.user)
		assert.Equal(r.T(), tc.expectedError, err)
		assert.Equal(r.T(), tc.expectedUserEntity, userEntity)
	}
}

func (r *RepositoryTestSuite) TestUpdate() {
	testCases := []struct {
		user          *entity.User
		ctx           context.Context
		expectedError error
	}{
		{
			user: &entity.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Password:  "pass",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			ctx:           context.Background(),
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		r.mock.ExpectExec("UPDATE users").
			WillReturnResult(sqlmock.NewResult(0, 1))
		err := userRepository.Update(tc.ctx, tc.user)
		assert.Equal(r.T(), tc.expectedError, err)
	}
}

func (r *RepositoryTestSuite) TestRemove() {
	testCases := []struct {
		id            int64
		ctx           context.Context
		expectedError error
	}{
		{
			id:            1,
			ctx:           context.Background(),
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		r.mock.ExpectExec("DELETE FROM users").
			WillReturnResult(sqlmock.NewResult(0, 1))
		err := userRepository.Remove(tc.ctx, tc.id)
		assert.Equal(r.T(), tc.expectedError, err)
	}
}

func (r *RepositoryTestSuite) TestGetByID() {
	testCases := []struct {
		id                 int64
		ctx                context.Context
		expectedUserEntity *entity.User
		expectedError      error
	}{
		{
			id:  1,
			ctx: context.Background(),
			expectedUserEntity: &entity.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Email:     "test@gmail.com",
				Country:   "UK",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		rows := r.mock.NewRows([]string{"id", "first_name", "last_name", "nick_name", "email", "country", "created_at", "updated_at"}).
			AddRow(
				tc.expectedUserEntity.ID,
				tc.expectedUserEntity.FirstName,
				tc.expectedUserEntity.LastName,
				tc.expectedUserEntity.NickName,
				tc.expectedUserEntity.Email,
				tc.expectedUserEntity.Country,
				tc.expectedUserEntity.CreatedAt,
				tc.expectedUserEntity.UpdatedAt,
			)

		r.mock.ExpectQuery("SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users").
			WithArgs(tc.id).
			WillReturnRows(rows)
		userEntity, err := userRepository.GetByID(tc.ctx, tc.id)
		assert.Equal(r.T(), tc.expectedError, err)
		assert.Equal(r.T(), tc.expectedUserEntity, userEntity)
	}
}

func (r *RepositoryTestSuite) TestGetByNickName() {
	testCases := []struct {
		nickname           string
		ctx                context.Context
		expectedUserEntity *entity.User
		expectedError      error
	}{
		{
			nickname: "test",
			ctx:      context.Background(),
			expectedUserEntity: &entity.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Email:     "test@gmail.com",
				Country:   "UK",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		rows := r.mock.NewRows([]string{"id", "first_name", "last_name", "nick_name", "email", "country", "created_at", "updated_at"}).
			AddRow(
				tc.expectedUserEntity.ID,
				tc.expectedUserEntity.FirstName,
				tc.expectedUserEntity.LastName,
				tc.expectedUserEntity.NickName,
				tc.expectedUserEntity.Email,
				tc.expectedUserEntity.Country,
				tc.expectedUserEntity.CreatedAt,
				tc.expectedUserEntity.UpdatedAt,
			)

		r.mock.ExpectQuery("SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users").
			WithArgs(tc.nickname).
			WillReturnRows(rows)
		userEntity, err := userRepository.GetByNickName(tc.ctx, tc.nickname)
		assert.Equal(r.T(), tc.expectedError, err)
		assert.Equal(r.T(), tc.expectedUserEntity, userEntity)
	}
}

func (r *RepositoryTestSuite) TestGetByEmail() {
	testCases := []struct {
		email              string
		ctx                context.Context
		expectedUserEntity *entity.User
		expectedError      error
	}{
		{
			email: "test@gmail.com",
			ctx:   context.Background(),
			expectedUserEntity: &entity.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Email:     "test@gmail.com",
				Country:   "UK",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			},
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		rows := r.mock.NewRows([]string{"id", "first_name", "last_name", "nick_name", "email", "country", "created_at", "updated_at"}).
			AddRow(
				tc.expectedUserEntity.ID,
				tc.expectedUserEntity.FirstName,
				tc.expectedUserEntity.LastName,
				tc.expectedUserEntity.NickName,
				tc.expectedUserEntity.Email,
				tc.expectedUserEntity.Country,
				tc.expectedUserEntity.CreatedAt,
				tc.expectedUserEntity.UpdatedAt,
			)

		r.mock.ExpectQuery("SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users").
			WithArgs(tc.email).
			WillReturnRows(rows)
		userEntity, err := userRepository.GetByEmail(tc.ctx, tc.email)
		assert.Equal(r.T(), tc.expectedError, err)
		assert.Equal(r.T(), tc.expectedUserEntity, userEntity)
	}
}

func (r *RepositoryTestSuite) TestGet() {
	testCases := []struct {
		filter               *entity.Filter
		ctx                  context.Context
		page                 int64
		pageSize             int64
		expectedUserEntities []*entity.User
		expectedError        error
	}{
		{
			filter: &entity.Filter{
				Country: "UK",
			},
			ctx:      context.Background(),
			page:     1,
			pageSize: 10,
			expectedUserEntities: []*entity.User{
				{
					ID:        1,
					FirstName: "test",
					LastName:  "test",
					NickName:  "test",
					Email:     "test@gmail.com",
					Country:   "UK",
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {

		rows := r.mock.NewRows([]string{"id", "first_name", "last_name", "nick_name", "email", "country", "created_at", "updated_at"})
		for _, expectedUserEntity := range tc.expectedUserEntities {
			rows.AddRow(
				expectedUserEntity.ID,
				expectedUserEntity.FirstName,
				expectedUserEntity.LastName,
				expectedUserEntity.NickName,
				expectedUserEntity.Email,
				expectedUserEntity.Country,
				expectedUserEntity.CreatedAt,
				expectedUserEntity.UpdatedAt,
			)
		}

		r.mock.ExpectQuery("SELECT id, first_name, last_name, nick_name, email, country, created_at, updated_at FROM users").
			WillReturnRows(rows)
		userEntities, err := userRepository.Get(tc.ctx, tc.filter, tc.page, tc.pageSize)
		assert.Equal(r.T(), tc.expectedError, err)
		assert.Equal(r.T(), tc.expectedUserEntities, userEntities)
	}
}

func (r *RepositoryTestSuite) TestGetCount() {
	testCases := []struct {
		filter        *entity.Filter
		ctx           context.Context
		expectedCount uint64
		expectedError error
	}{
		{
			filter: &entity.Filter{
				Country: "UK",
			},
			ctx:           context.Background(),
			expectedCount: 1,
			expectedError: nil,
		},
	}

	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {

		rows := r.mock.NewRows([]string{"total"}).AddRow(tc.expectedCount)

		r.mock.ExpectQuery("SELECT count\\(\\*\\) as total FROM users").
			WillReturnRows(rows)
		count, err := userRepository.GetCount(tc.ctx, tc.filter)
		assert.Equal(r.T(), tc.expectedError, err)
		assert.Equal(r.T(), tc.expectedCount, count)
	}
}

func (r *RepositoryTestSuite) TestPublishUserChangeEvent() {
	testCases := []struct {
		ID int64
	}{
		{ID: 1},
		{ID: 2},
	}
	r.db, r.mock = databaseMocks.NewDBMock()
	r.redis = redisMocks.NewRedisMock()
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{r.redis.Addr()},
	})
	userRepository := NewUserRepository(r.db, redisClient)

	for _, tc := range testCases {
		err := userRepository.PublishUserChangeEvent(tc.ID)
		assert.Nil(r.T(), err)
	}

}

func TestRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RepositoryTestSuite))
}
