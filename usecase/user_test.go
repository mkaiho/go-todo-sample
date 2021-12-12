package usecase

import (
	"errors"
	"testing"

	"github.com/mkaiho/go-todo-sample/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewGetUsersUseCase(t *testing.T) {
	type args struct {
		userDataAccess UsersDataAccess
	}
	tests := []struct {
		name string
		args args
		want GetUsersUseCase
	}{
		{
			name: "Return GetUsersUseCase",
			args: args{
				userDataAccess: new(MockUsersDataAccess),
			},
			want: &getUsersInteractor{
				userDataAccess: new(MockUsersDataAccess),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewGetUsersUseCase(tt.args.userDataAccess)
			assert.Equal(t, tt.want, got, "NewGetUsersUseCase() = %v, want %v", got, tt.want)
		})
	}
}

func TestGetUsersUseCaseInput_toInvalidError(t *testing.T) {
	type fields struct {
		ID entity.UserID
	}
	tests := []struct {
		name   string
		fields fields
		want   UseCaseError
	}{
		{
			name: "Return Invalid Usecase Input Error",
			fields: fields{
				ID: func() entity.UserID {
					id := entity.MockUserID{}
					id.On("String").Return("dummy")
					return &id
				}(),
			},
			want: &useCaseError{
				code:    ErrCodeInvalidUsecaseInput,
				message: "GetUsersUseCase input is invalid: map[id:dummy]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := GetUsersUseCaseInput{
				ID: tt.fields.ID,
			}
			got := i.toInvalidError()
			assert.Equal(t, tt.want, got, "GetUsersUseCaseInput.toInvalidError() = %v, want %v", got, tt.want)
		})
	}
}

func Test_getUsersInteractor_Get(t *testing.T) {
	type fields struct {
		userDataAccess UsersDataAccess
	}
	type args struct {
		input GetUsersUseCaseInput
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return user when input is valid and user exists",
			fields: fields{
				userDataAccess: func() UsersDataAccess {
					dataAccess := new(MockUsersDataAccess)
					dataAccess.On("FindByID", mock.Anything).Return(new(entity.MockUser), nil)
					return dataAccess
				}(),
			},
			args: args{
				GetUsersUseCaseInput{
					ID: func() entity.UserID {
						id := new(entity.MockUserID)
						id.On("Validate").Return(nil)
						return id
					}(),
				},
			},
			want:    new(entity.MockUser),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when input.ID is invalid",
			fields: fields{
				userDataAccess: func() UsersDataAccess {
					dataAccess := new(MockUsersDataAccess)
					dataAccess.On("FindByID", mock.Anything).Return(new(entity.MockUser), nil)
					return dataAccess
				}(),
			},
			args: args{
				GetUsersUseCaseInput{
					ID: func() entity.UserID {
						id := new(entity.MockUserID)
						id.On("String").Return("test_id_dummy")
						id.On("Validate").Return(errors.New("dummy error"))
						return id
					}(),
				},
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				id := new(entity.MockUserID)
				id.On("String").Return("test_id_dummy")
				return assert.Equal(t, e, NewErrInvalidUsecaseInput("GetUsersUseCase", map[string]interface{}{
					"id": id,
				}), "getUsersInteractor.Get() error = %v, wantErr %v", e)
			},
		},
		{
			name: "Return error when UsersDataAccess faild to find user",
			fields: fields{
				userDataAccess: func() UsersDataAccess {
					dataAccess := new(MockUsersDataAccess)
					dataAccess.On("FindByID", mock.Anything).Return(nil, &useCaseError{code: 999, message: "dummy error"})
					return dataAccess
				}(),
			},
			args: args{
				GetUsersUseCaseInput{
					ID: func() entity.UserID {
						id := new(entity.MockUserID)
						id.On("Validate").Return(nil)
						return id
					}(),
				},
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "dummy error")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &getUsersInteractor{
				userDataAccess: tt.fields.userDataAccess,
			}
			got, err := i.Get(tt.args.input)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "getUsersInteractor.Get() = %v, want %v", got, tt.want)
		})
	}
}
