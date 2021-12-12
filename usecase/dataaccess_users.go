package usecase

import "github.com/mkaiho/go-todo-sample/entity"

type UsersDataAccess interface {
	FindByID(id entity.UserID) (entity.User, UseCaseError)
}
