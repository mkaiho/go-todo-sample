package rdb

import (
	"errors"
	"testing"

	"github.com/mkaiho/go-todo-sample/entity"
	"github.com/mkaiho/go-todo-sample/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func NewUserMock(
	id string, invalidIDErr error,
	name string, invalidNameErr error,
	email string, invalidEmailErr error,
) entity.User {
	mockID := new(entity.MockUserID)
	mockID.On("String").Return(id)
	mockID.On("IsEmpty").Return(len(id))
	mockID.On("IsValid").Return(invalidIDErr)
	mockName := new(entity.MockUserName)
	mockName.On("String").Return(name)
	mockName.On("IsEmpty").Return(len(name))
	mockName.On("IsValid").Return(invalidNameErr)
	mockEmail := new(entity.MockEmail)
	mockEmail.On("String").Return(email)
	mockEmail.On("IsEmpty").Return(len(email))
	mockEmail.On("IsValid").Return(invalidEmailErr)
	mockUserer := new(entity.MockUser)
	mockUserer.On("UserID").Return(mockID)
	mockUserer.On("Name").Return(mockName)
	mockUserer.On("Email").Return(mockEmail)
	return mockUserer
}

func Test_userRow_toUserEntity(t *testing.T) {
	createMockUser := func(
		id string, invalidIDErr usecase.UseCaseError,
		name string, invalidNameErr usecase.UseCaseError,
		email string, invalidEmailErr usecase.UseCaseError,
	) entity.User {
		mockID := new(entity.MockUserID)
		mockID.On("String").Return(id)
		mockID.On("IsEmpty").Return(len(id))
		mockID.On("IsValid").Return(invalidIDErr)
		mockName := new(entity.MockUserName)
		mockName.On("String").Return(name)
		mockName.On("IsEmpty").Return(len(name))
		mockName.On("IsValid").Return(invalidNameErr)
		mockEmail := new(entity.MockEmail)
		mockEmail.On("String").Return(email)
		mockEmail.On("IsEmpty").Return(len(email))
		mockEmail.On("IsValid").Return(invalidEmailErr)
		mockUserer := new(entity.MockUser)
		mockUserer.On("UserID").Return(mockID)
		mockUserer.On("Name").Return(mockName)
		mockUserer.On("Email").Return(mockEmail)
		return mockUserer
	}
	type fields struct {
		ID    string
		Name  string
		Email string
	}
	tests := []struct {
		name    string
		fields  fields
		want    entity.User
		wantErr func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool
	}{
		{
			name: "Return User from userRow",
			fields: fields{
				ID:    "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c",
				Name:  "test_user",
				Email: "xxx.yyy@zzz.com",
			},
			want: createMockUser("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c", nil, "test_user", nil, "xxx.yyy@zzz.com", nil),
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				return assert.Nil(tt, e)
			},
		},
		{
			name: "Return error when id is invalid",
			fields: fields{
				ID:    "dummy_id",
				Name:  "test_user",
				Email: "xxx.yyy@zzz.com",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(t, e.Code(), wantCode, "userRow.toUserEntity() error.Code() got = %v, want %v", e.Code(), wantCode) {
					return false
				}
				wantErr := "unknown error is occured: id format is invalid: \"dummy_id\""
				return !assert.Equal(t, wantErr, e.Error(), "userRow.toUserEntity() error.Error() got = %v, want %v", e.Error(), wantErr)
			},
		},
		{
			name: "Return error when name is invalid",
			fields: fields{
				ID:    "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c",
				Name:  "",
				Email: "xxx.yyy@zzz.com",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(t, e.Code(), wantCode, "userRow.toUserEntity() error.Code() got = %v, want %v", e.Code(), wantCode) {
					return false
				}
				wantErr := "unknown error is occured: invalid user name"
				return !assert.Equal(t, wantErr, e.Error(), "userRow.toUserEntity() error.Error() got = %v, want %v", e.Error(), wantErr)
			},
		},
		{
			name: "Return error when email is invalid",
			fields: fields{
				ID:    "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c",
				Name:  "test_user",
				Email: "xxx.yyy@",
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(t, e.Code(), wantCode, "userRow.toUserEntity() error.Code() got = %v, want %v", e.Code(), wantCode) {
					return false
				}
				wantErr := "unknown error is occured: email format is invalid: \"xxx.yyy@\""
				return !assert.Equal(t, wantErr, e.Error(), "userRow.toUserEntity() error.Error() got = %v, want %v", e.Error(), wantErr)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &userRow{
				ID:    tt.fields.ID,
				Name:  tt.fields.Name,
				Email: tt.fields.Email,
			}
			got, err := r.toUserEntity()
			if !tt.wantErr(t, err) {
				return
			}
			wantUserID := tt.want.UserID().String()
			actualUserID := got.UserID().String()
			if !assert.Equal(t, wantUserID, actualUserID, "userRow.toUserEntity().UserID().String() got = %v, want %v", actualUserID, wantUserID) {
				return
			}
			wantName := tt.want.Name().String()
			actualName := got.Name().String()
			if !assert.Equal(t, wantName, actualName, "userRow.toUserEntity().Name().String() got = %v, want %v", actualName, wantName) {
				return
			}
			wantEmail := tt.want.Email().String()
			actualEmail := got.Email().String()
			if !assert.Equal(t, wantEmail, actualEmail, "userRow.toUserEntity().Email().String() got = %v, want %v", actualEmail, wantName) {
				return
			}
		})
	}
}

func TestNewUsersDataAccess(t *testing.T) {
	type args struct {
		idgen IDGenerator
		ds    DataSource
	}
	tests := []struct {
		name string
		args args
		want usecase.UsersDataAccess
	}{
		{
			name: "Return UsersDataAccess",
			args: args{
				idgen: &MockIDGenerator{},
				ds:    &MockDataSource{},
			},
			want: &usersDataAccess{
				idgen: &MockIDGenerator{},
				ds:    &MockDataSource{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewUsersDataAccess(tt.args.idgen, tt.args.ds)
			assert.Equal(t, tt.want, got, "NewUsersDataAccess() = %v, want %v", got, tt.want)
		})
	}
}

func Test_usersDataAccess_FindByID(t *testing.T) {
	type fields struct {
		idgen IDGenerator
		ds    DataSource
	}
	type args struct {
		id entity.UserID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    entity.User
		wantErr func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool
	}{
		{
			name: "Return User",
			fields: fields{
				idgen: &MockIDGenerator{},
				ds: func() DataSource {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					mockDataSource := new(MockDataSource)
					mockDataSource.On("Query", mock.Anything, &query{
						columnNames:    []string{"id", "name", "email"},
						tableName:      "users",
						joinQuery:      "",
						joinQueryArgs:  []interface{}{},
						whereQuery:     "id = ?",
						whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
					}).Return(nil).Once().Run(func(args mock.Arguments) {
						row := args[0].(*userRow)
						row.ID = "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"
						row.Name = "test_user"
						row.Email = "xxx.yyy@zzz.com"
					})
					return mockDataSource
				}(),
			},
			args: args{
				id: func() entity.UserID {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					return mockID
				}(),
			},
			want: NewUserMock("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c", nil, "test_user", nil, "xxx.yyy@zzz.com", nil),
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				return assert.Nil(tt, e)
			},
		},
		{
			name: "Return error when failed to query",
			fields: fields{
				idgen: &MockIDGenerator{},
				ds: func() DataSource {
					mockDataSource := new(MockDataSource)
					mockDataSource.On("Query", mock.Anything, mock.Anything).Return(usecase.NewErrUnknown(errors.New("dummy error")))
					return mockDataSource
				}(),
			},
			args: args{
				id: func() entity.UserID {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					return mockID
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(tt, wantCode, e.Code(), "usersDataAccess.FindByID() got.Code() = %v, want.Code() %v", e.Code(), wantCode) {
					return false
				}
				wantError := "unknown error is occured: dummy error"
				return assert.Equal(tt, wantError, e.Error(), "usersDataAccess.FindByID() got.Error() = %v, want.Error() %v", e.Error(), wantError)
			},
		},
		{
			name: "Return error when id is invalid",
			fields: fields{
				idgen: &MockIDGenerator{},
				ds: func() DataSource {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("xxxxx")
					mockDataSource := new(MockDataSource)
					mockDataSource.On("Query", mock.Anything, &query{
						columnNames:    []string{"id", "name", "email"},
						tableName:      "users",
						joinQuery:      "",
						joinQueryArgs:  []interface{}{},
						whereQuery:     "id = ?",
						whereQueryArgs: []interface{}{"xxxxx"},
					}).Return(nil).Once().Run(func(args mock.Arguments) {
						row := args[0].(*userRow)
						row.ID = "xxxxx"
						row.Name = "test_user"
						row.Email = "xxx.yyy@zzz.com"
					})
					return mockDataSource
				}(),
			},
			args: args{
				id: func() entity.UserID {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("xxxxx")
					return mockID
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(tt, wantCode, e.Code(), "usersDataAccess.FindByID() got.Code() = %v, want.Code() %v", e.Code(), wantCode) {
					return false
				}
				wantError := "unknown error is occured: id format is invalid: \"xxxxx\""
				return assert.Equal(tt, wantError, e.Error(), "usersDataAccess.FindByID() got.Error() = %v, want.Error() %v", e.Error(), wantError)
			},
		},
		{
			name: "Return error when name is invalid",
			fields: fields{
				idgen: &MockIDGenerator{},
				ds: func() DataSource {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					mockDataSource := new(MockDataSource)
					mockDataSource.On("Query", mock.Anything, &query{
						columnNames:    []string{"id", "name", "email"},
						tableName:      "users",
						joinQuery:      "",
						joinQueryArgs:  []interface{}{},
						whereQuery:     "id = ?",
						whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
					}).Return(nil).Once().Run(func(args mock.Arguments) {
						row := args[0].(*userRow)
						row.ID = "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"
						row.Name = ""
						row.Email = "xxx.yyy@zzz.com"
					})
					return mockDataSource
				}(),
			},
			args: args{
				id: func() entity.UserID {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					return mockID
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(tt, wantCode, e.Code(), "usersDataAccess.FindByID() got.Code() = %v, want.Code() %v", e.Code(), wantCode) {
					return false
				}
				wantError := "unknown error is occured: invalid user name"
				return assert.Equal(tt, wantError, e.Error(), "usersDataAccess.FindByID() got.Error() = %v, want.Error() %v", e.Error(), wantError)
			},
		},
		{
			name: "Return error when email is invalid",
			fields: fields{
				idgen: &MockIDGenerator{},
				ds: func() DataSource {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					mockDataSource := new(MockDataSource)
					mockDataSource.On("Query", mock.Anything, &query{
						columnNames:    []string{"id", "name", "email"},
						tableName:      "users",
						joinQuery:      "",
						joinQueryArgs:  []interface{}{},
						whereQuery:     "id = ?",
						whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
					}).Return(nil).Once().Run(func(args mock.Arguments) {
						row := args[0].(*userRow)
						row.ID = "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"
						row.Name = "test_user"
						row.Email = "xxx.yyy@"
					})
					return mockDataSource
				}(),
			},
			args: args{
				id: func() entity.UserID {
					mockID := new(entity.MockUserID)
					mockID.On("String").Return("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c")
					return mockID
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(tt, wantCode, e.Code(), "usersDataAccess.FindByID() got.Code() = %v, want.Code() %v", e.Code(), wantCode) {
					return false
				}
				wantError := "unknown error is occured: email format is invalid: \"xxx.yyy@\""
				return assert.Equal(tt, wantError, e.Error(), "usersDataAccess.FindByID() got.Error() = %v, want.Error() %v", e.Error(), wantError)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			da := &usersDataAccess{
				idgen: tt.fields.idgen,
				ds:    tt.fields.ds,
			}
			got, err := da.FindByID(tt.args.id)
			if !tt.wantErr(t, err) {
				return
			}
			if tt.want == nil {
				assert.Equal(t, tt.want, got, "usersDataAccess.FindByID() got = %v, want %v", got, tt.want)
				return
			}
			if gotUserID, wantUserID := got.UserID().String(), tt.want.UserID().String(); !assert.Equal(t, wantUserID, gotUserID,
				"usersDataAccess.FindByID() got.UserID().String() = %v, want.UserID().String() %v", gotUserID, wantUserID,
			) {
				return
			}
			if gotUserName, wantUserName := got.Name().String(), tt.want.Name().String(); !assert.Equal(t, wantUserName, gotUserName,
				"usersDataAccess.FindByID() got.Name().String() = %v, want.Name().String() %v", gotUserName, wantUserName,
			) {
				return
			}
			if gotUserEmail, wantUserEmail := got.Email().String(), tt.want.Email().String(); !assert.Equal(t, wantUserEmail, gotUserEmail,
				"usersDataAccess.FindByID() got.Email().String() = %v, want.Email().String() %v", gotUserEmail, wantUserEmail,
			) {
				return
			}
		})
	}
}
