package utils

import (
	"faceit/domain/user/dto"
	"faceit/domain/user/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ConvertorTestSuite struct {
	suite.Suite
}

func (u *ConvertorTestSuite) TestUserDTOFromEntity() {
	userEntity := &entity.User{
		FirstName: "test",
		LastName:  "test",
		NickName:  "test",
		Email:     "test@gmail.com",
		Country:   "UK",
	}

	expectedUserDTO := &dto.User{
		FirstName: "test",
		LastName:  "test",
		NickName:  "test",
		Email:     "test@gmail.com",
		Country:   "UK",
	}
	userDTO := UserDTOFromEntity(userEntity)

	assert.Equal(u.T(), expectedUserDTO, userDTO)
}

func (u *ConvertorTestSuite) UserEntityFromDTO() {
	userDTO := &dto.User{
		FirstName: "test",
		LastName:  "test",
		NickName:  "test",
		Email:     "test@gmail.com",
		Country:   "UK",
	}

	expectedUserEntity := &entity.User{
		FirstName: "test",
		LastName:  "test",
		NickName:  "test",
		Email:     "test@gmail.com",
		Country:   "UK",
	}

	userEntity := UserEntityFromDTO(userDTO)

	assert.Equal(u.T(), expectedUserEntity, userEntity)
}

func TestConvertorTestSuite(t *testing.T) {
	suite.Run(t, new(ConvertorTestSuite))
}
