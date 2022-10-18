package service

import (
	"context"
	"faceit/domain/constants"
	"faceit/domain/user/dto"
	"faceit/domain/user/entity"
	mocks "faceit/mocks/domain/user/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type ServiceTestSuite struct {
	suite.Suite
}

func (s *ServiceTestSuite) TestCreate() {
	testCases := []struct {
		userEntity         *entity.User
		userDTO            *dto.User
		password           string
		expectedUserEntity *entity.User
		expectedUserDTO    *dto.User
		expectedError      error
	}{
		{
			userEntity: &entity.User{
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Password:  "pass",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			userDTO: &dto.User{
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			password: "pass",
			expectedUserEntity: &entity.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Password:  "pass",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			expectedUserDTO: &dto.User{
				ID:        1,
				FirstName: "test",
				LastName:  "test",
				NickName:  "test",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			expectedError: nil,
		},
	}

	repositoryMock := mocks.IUsersRepository{}
	for _, tc := range testCases {
		repositoryMock.On("Create", mock.Anything, tc.userEntity).Return(tc.expectedUserEntity, nil)
		repositoryMock.On("GetByEmail", mock.Anything, tc.userEntity.Email).Return(nil, constants.ErrUserNotFound)
		repositoryMock.On("GetByNickName", mock.Anything, tc.userEntity.NickName).Return(nil, constants.ErrUserNotFound)

		userService := NewUserService(&repositoryMock)
		userDTO, err := userService.Create(context.Background(), tc.userDTO, tc.password)
		assert.Equal(s.T(), tc.expectedError, err)
		assert.Equal(s.T(), tc.expectedUserDTO, userDTO)
	}
}

func (s *ServiceTestSuite) TestUpdate() {
	testCases := []struct {
		userEntity         *entity.User
		userDTO            *dto.User
		password           string
		expectedUserEntity *entity.User
		expectedError      error
	}{
		{
			userEntity: &entity.User{
				ID:        1,
				FirstName: "test2",
				LastName:  "test",
				NickName:  "test",
				Password:  "pass",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			userDTO: &dto.User{
				ID:        1,
				FirstName: "test2",
				LastName:  "test",
				NickName:  "test",
				Email:     "test@gmail.com",
				Country:   "UK",
			},
			password: "pass",
			expectedUserEntity: &entity.User{
				ID:        1,
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

	repositoryMock := mocks.IUsersRepository{}
	for _, tc := range testCases {
		repositoryMock.On("Update", mock.Anything, tc.userEntity).Return(nil)
		repositoryMock.On("GetByID", mock.Anything, tc.userEntity.ID).Return(tc.expectedUserEntity, nil)

		userService := NewUserService(&repositoryMock)
		err := userService.Update(context.Background(), tc.userDTO, tc.password)
		assert.Equal(s.T(), tc.expectedError, err)
	}
}

func (s *ServiceTestSuite) TestRemove() {
	testCases := []struct {
		id            int64
		expectedError error
	}{
		{
			id:            1,
			expectedError: nil,
		},
	}

	repositoryMock := mocks.IUsersRepository{}
	for _, tc := range testCases {
		repositoryMock.On("Remove", mock.Anything, tc.id).Return(nil)

		userService := NewUserService(&repositoryMock)
		err := userService.Remove(context.Background(), tc.id)
		assert.Equal(s.T(), tc.expectedError, err)
	}
}

func (s *ServiceTestSuite) TestGet() {
	testCases := []struct {
		filter               *dto.Filter
		entityFilter         *entity.Filter
		page                 int64
		pageSize             int64
		expectedCount        uint64
		expectedUserDTOs     []*dto.User
		expectedUserEntities []*entity.User
		expectedError        error
	}{
		{
			filter: &dto.Filter{
				Country: "UK",
			},
			entityFilter: &entity.Filter{
				Country: "UK",
			},
			page:          0,
			pageSize:      10,
			expectedCount: 1,
			expectedUserDTOs: []*dto.User{
				{
					ID:        1,
					FirstName: "test",
					LastName:  "test",
					NickName:  "test",
					Email:     "test@gmail.com",
					Country:   "UK",
				},
			},
			expectedUserEntities: []*entity.User{
				{
					ID:        1,
					FirstName: "test",
					LastName:  "test",
					NickName:  "test",
					Password:  "pass",
					Email:     "test@gmail.com",
					Country:   "UK",
				},
			},
			expectedError: nil,
		},
	}

	repositoryMock := mocks.IUsersRepository{}
	for _, tc := range testCases {
		repositoryMock.On("Get", mock.Anything, tc.entityFilter, tc.page, tc.pageSize).Return(tc.expectedUserEntities, tc.expectedError)
		repositoryMock.On("GetCount", mock.Anything, tc.entityFilter).Return(tc.expectedCount, tc.expectedError)

		userService := NewUserService(&repositoryMock)
		userDTOs, count, err := userService.Get(context.Background(), tc.filter, tc.page, tc.pageSize)
		assert.Equal(s.T(), tc.expectedError, err)
		assert.Equal(s.T(), tc.expectedCount, count)
		assert.Equal(s.T(), tc.expectedUserDTOs, userDTOs)
	}
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
