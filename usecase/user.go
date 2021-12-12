package usecase

import (
	"github.com/mkaiho/go-todo-sample/entity"
)

/** List users use case **/
var _ GetUsersUseCase = (*getUsersInteractor)(nil)

func NewGetUsersUseCase(userDataAccess UsersDataAccess) GetUsersUseCase {
	return &getUsersInteractor{
		userDataAccess: userDataAccess,
	}
}

type GetUsersUseCaseInput struct {
	ID entity.UserID
}

func (i GetUsersUseCaseInput) toInvalidError() UseCaseError {
	const name = "GetUsersUseCase"
	return NewErrInvalidUsecaseInput(name, map[string]interface{}{
		"id": i.ID,
	})
}

type GetUsersUseCase interface {
	Get(input GetUsersUseCaseInput) (entity.User, UseCaseError)
}

type getUsersInteractor struct {
	userDataAccess UsersDataAccess
}

func (i *getUsersInteractor) Get(input GetUsersUseCaseInput) (entity.User, UseCaseError) {
	if err := input.ID.Validate(); err != nil {
		return nil, input.toInvalidError()
	}
	user, err := i.userDataAccess.FindByID(input.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}
