package service

import (
	"context"
	"faceit/domain/dto"
	"faceit/domain/entity"
	"faceit/domain/user/repository"
	"faceit/domain/utils"
)

type IUserService interface {
	Create(ctx context.Context, user *dto.User, password string) (*dto.User, error)
	Update(ctx context.Context, user *dto.User, password string) error
	Remove(ctx context.Context, id int64) error
	Get(ctx context.Context, filter *dto.Filter, page, pageSize int64) ([]*dto.User, int64, error)
}

type UserService struct {
	repository repository.IUsersRepository
}

func NewUserService(repository repository.IUsersRepository) *UserService {
	return &UserService{repository: repository}
}

func (u *UserService) Create(ctx context.Context, user *dto.User, password string) (*dto.User, error) {
	userEntity := utils.UserEntityFromDTO(user)
	userEntity.Password = password
	createdUserEntity, err := u.repository.Create(ctx, userEntity)
	if err != nil {
		return nil, err
	}

	return utils.UserDTOFromEntity(createdUserEntity), nil
}

func (u *UserService) Update(ctx context.Context, user *dto.User, password string) error {
	userEntity := utils.UserEntityFromDTO(user)
	userEntity.Password = password
	err := u.repository.Update(ctx, userEntity)
	return err
}

func (u *UserService) Remove(ctx context.Context, id int64) error {
	err := u.repository.Remove(ctx, id)
	return err
}

func (u *UserService) Get(ctx context.Context, filter *dto.Filter, page, pageSize int64) ([]*dto.User, int64, error) {
	filterEntity := entity.FilterEntityFromDTO(filter)
	userEntities, count, err := u.repository.Get(ctx, filterEntity, page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	userDTOs := make([]*dto.User, len(userEntities))

	for i, userEntity := range userEntities {
		userDTOs[i] = utils.UserDTOFromEntity(userEntity)
	}

	return userDTOs, count, nil
}
