package rdb

import (
	"github.com/mkaiho/go-todo-sample/entity"
	"github.com/mkaiho/go-todo-sample/usecase"
)

/** row **/
type userRow struct {
	ID    string
	Name  string
	Email string
}

func (r *userRow) toUserEntity() (entity.User, usecase.UseCaseError) {
	id, err := entity.ParseUserID(r.ID)
	if err != nil {
		return nil, usecase.NewErrUnknown(err)
	}
	name, err := entity.ParseUserName(r.Name)
	if err != nil {
		return nil, usecase.NewErrUnknown(err)
	}
	email, err := entity.ParseEmail(r.Email)
	if err != nil {
		return nil, usecase.NewErrUnknown(err)
	}

	return entity.NewUser(entity.UserID(id), name, email), nil
}

/** UsersDataAccess **/
func NewUsersDataAccess(idgen IDGenerator, ds DataSource) usecase.UsersDataAccess {
	return &usersDataAccess{
		idgen: idgen,
		ds:    ds,
	}
}

type usersDataAccess struct {
	idgen IDGenerator
	ds    DataSource
}

func (da *usersDataAccess) FindByID(id entity.UserID) (entity.User, usecase.UseCaseError) {
	query, err := NewQueryBuilder().Select("id", "name", "email").From("users").Where("id = ?", id.String()).Build()
	if err != nil {
		return nil, err
	}

	row := userRow{}
	err = da.ds.Query(&row, query)
	if err != nil {
		return nil, err
	}

	user, err := row.toUserEntity()
	if err != nil {
		return nil, err
	}

	return user, nil
}
