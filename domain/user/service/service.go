package service

import (
	"context"
	"errors"
	"faceit/domain/constants"
	"faceit/domain/user/dto"
	"faceit/domain/user/entity"
	"faceit/domain/user/repository"
	"faceit/domain/user/utils"
	"strings"
)

type IUserService interface {
	Create(ctx context.Context, user *dto.User, password string) (*dto.User, error)
	Update(ctx context.Context, user *dto.User, password string) error
	Remove(ctx context.Context, id int64) error
	Get(ctx context.Context, filter *dto.Filter, page, pageSize int64) ([]*dto.User, uint64, error)
}

type UserService struct {
	repository repository.IUsersRepository
}

func NewUserService(repository repository.IUsersRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) Create(ctx context.Context, user *dto.User, password string) (*dto.User, error) {
	// check for email and nickname uniqueness
	foundUserEntity, err := u.repository.GetByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, constants.ErrUserNotFound) {
		return nil, err
	}
	if foundUserEntity != nil {
		return nil, constants.ErrUserExists
	}

	foundUserEntity, err = u.repository.GetByNickName(ctx, user.NickName)
	if err != nil && !errors.Is(err, constants.ErrUserNotFound) {
		return nil, err
	}
	if foundUserEntity != nil {
		return nil, constants.ErrUserExists
	}

	userEntity := utils.UserEntityFromDTO(user)
	userEntity.Password = password
	// convert the user's country to uppercase for consistency
	userEntity.Country = strings.ToUpper(userEntity.Country)
	createdUserEntity, err := u.repository.Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return utils.UserDTOFromEntity(createdUserEntity), nil
}

func (u *UserService) Update(ctx context.Context, user *dto.User, password string) error {
	// check if the user exists
	foundUserEntity, err := u.repository.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}

	// check if there are any changes, if not return error
	var hasChanges bool
	if user.FirstName != "" && user.FirstName != foundUserEntity.FirstName {
		hasChanges = true
	}
	if user.LastName != "" && user.LastName != foundUserEntity.LastName {
		hasChanges = true
	}
	if user.NickName != "" && user.NickName != foundUserEntity.NickName {
		hasChanges = true
	}
	if user.Email != "" && user.Email != foundUserEntity.Email {
		hasChanges = true
	}
	if user.Country != "" && user.Country != foundUserEntity.Country {
		hasChanges = true
	}
	if password != "" && password != foundUserEntity.Password {
		hasChanges = true
	}
	if !hasChanges {
		return constants.ErrHasNoChanges
	}

	userEntity := utils.UserEntityFromDTO(user)
	userEntity.Password = password
	err = u.repository.Update(ctx, userEntity)
	return err
}

func (u *UserService) Remove(ctx context.Context, id int64) error {
	err := u.repository.Remove(ctx, id)
	return err
}

func (u *UserService) Get(ctx context.Context, filter *dto.Filter, page, pageSize int64) ([]*dto.User, uint64, error) {
	filterEntity := entity.FilterEntityFromDTO(filter)
	userEntities, err := u.repository.Get(ctx, filterEntity, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	userDTOs := make([]*dto.User, len(userEntities))

	for i, userEntity := range userEntities {
		userDTOs[i] = utils.UserDTOFromEntity(userEntity)
	}

	// get the total count of the users for pagination
	count, err := u.repository.GetCount(ctx, filterEntity)
	if err != nil {
		return nil, 0, err
	}

	return userDTOs, count, nil
}
