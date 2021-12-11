package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseEmail(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    email
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return email",
			args: args{
				value: "xxx.yyy@zzz.com",
			},
			want:    email("xxx.yyy@zzz.com"),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when email format is invalid",
			args: args{
				value: "xxx.yyyzzz.com",
			},
			want: "",
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "email format is invalid: \"xxx.yyyzzz.com\"")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseEmail(tt.args.value)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "ParseEmail() = %v, want %v", got, tt.want)
		})
	}
}

func Test_email_String(t *testing.T) {
	tests := []struct {
		name string
		em   email
		want string
	}{
		{
			name: "Return value converted string",
			em:   email("xxx.yyy@zzz.com"),
			want: "xxx.yyy@zzz.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.em.String()
			assert.Equal(t, tt.want, got, "email.String() = %v, want %v", got, tt.want)
		})
	}
}

func Test_email_Validate(t *testing.T) {
	tests := []struct {
		name    string
		em      email
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Return no error when passed \"xxx.yyy@zzz.com\"",
			em:      email("xxx.yyy@zzz.com"),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when Email is empty",
			em:   email(""),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "email is empty", "email.Validate() error = %v", e)
			},
		},
		{
			name: "Return error when passed value does not contain \"@\"",
			em:   email("xxx.yyyzzz.com"),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "email format is invalid: \"xxx.yyyzzz.com\"", "email.Validate() error = %v", e)
			},
		},
		{
			name: "Return error when passed value starts with \"@\"",
			em:   email("@zzz.com"),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "email format is invalid: \"@zzz.com\"", "email.Validate() error = %v", e)
			},
		},
		{
			name: "Return error when passed value ends with \"@\"",
			em:   email("xxx.yyy@"),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "email format is invalid: \"xxx.yyy@\"", "email.Validate() error = %v", e)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.em.Validate(); !tt.wantErr(t, err, "email.Validate() error = %v, wantErr %v", err) {
				return
			}
		})
	}
}

func TestNewUser(t *testing.T) {
	type args struct {
		name  UserName
		email Email
	}
	type mocks struct {
		UserID UserID
	}
	tests := []struct {
		name    string
		args    args
		mocks   mocks
		want    User
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return User",
			args: args{
				name:  userName("testuser"),
				email: email("xxx.yyy@zzz.com"),
			},
			mocks: mocks{
				UserID: func() UserID {
					m := new(MockUserID)
					m.On("IsEmpty").Return(false)
					m.On("Validate").Return(nil)
					m.On("String").Return("id_xxxxx")
					return m
				}(),
			},
			want: &user{
				name:  userName("testuser"),
				email: email("xxx.yyy@zzz.com"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return User when user id is empty",
			args: args{
				name:  userName("testuser"),
				email: email("xxx.yyy@zzz.com"),
			},
			mocks: mocks{
				UserID: func() UserID {
					m := new(MockUserID)
					m.On("IsEmpty").Return(true)
					m.On("Validate").Return(errors.New("dummy error"))
					m.On("String").Return("")
					return m
				}(),
			},
			want: &user{
				name:  userName("testuser"),
				email: email("xxx.yyy@zzz.com"),
			},
			wantErr: assert.NoError,
		},
		{
			name: "Return error when user id is invalid",
			args: args{
				name:  userName("testuser"),
				email: email("xxx.yyy@zzz.com"),
			},
			mocks: mocks{
				UserID: func() UserID {
					m := new(MockUserID)
					m.On("IsEmpty").Return(false)
					m.On("Validate").Return(errors.New("dummy error"))
					m.On("String").Return("")
					return m
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "dummy error")
			},
		},
		{
			name: "Return error when user name is invalid",
			args: args{
				name:  userName(""),
				email: email("xxx.yyy@zzz.com"),
			},
			mocks: mocks{
				UserID: func() UserID {
					m := new(MockUserID)
					m.On("IsEmpty").Return(false)
					m.On("Validate").Return(nil)
					m.On("String").Return("id_xxxxx")
					return m
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "invalid user name")
			},
		},
		{
			name: "Return error when email format is invalid",
			args: args{
				name:  userName("testuser"),
				email: email("xxx.invalid"),
			},
			mocks: mocks{
				UserID: func() UserID {
					m := new(MockUserID)
					m.On("IsEmpty").Return(false)
					m.On("Validate").Return(nil)
					m.On("String").Return("id_xxxxx")
					return m
				}(),
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "email format is invalid: \"xxx.invalid\"")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id := tt.mocks.UserID
			got, err := NewUser(id, tt.args.name, tt.args.email)
			if !tt.wantErr(t, err) {
				return
			}
			if tt.want == nil && assert.Nil(t, got) {
				return
			}
			if !assert.NotNil(t, got) {
				return
			}
			assert.Equal(t, id, got.UserID())
			assert.Equal(t, tt.want.Name(), got.Name())
			assert.Equal(t, tt.want.Email(), got.Email())
		})
	}
}

func Test_user_UserID(t *testing.T) {
	type fields struct {
		id    UserID
		name  UserName
		email Email
	}
	tests := []struct {
		name   string
		fields fields
		want   UserID
	}{
		{
			name: "Return UserID",
			fields: fields{
				id: new(MockUserID),
			},
			want: new(MockUserID),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:    tt.fields.id,
				name:  tt.fields.name,
				email: tt.fields.email,
			}
			got := u.UserID()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_user_Name(t *testing.T) {
	type fields struct {
		id    UserID
		name  UserName
		email Email
	}
	tests := []struct {
		name   string
		fields fields
		want   UserName
	}{
		{
			name: "Return UserID",
			fields: fields{
				name: new(MockUserName),
			},
			want: new(MockUserName),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:    tt.fields.id,
				name:  tt.fields.name,
				email: tt.fields.email,
			}
			got := u.Name()
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_user_Email(t *testing.T) {
	type fields struct {
		id    UserID
		name  UserName
		email Email
	}
	tests := []struct {
		name   string
		fields fields
		want   Email
	}{
		{
			name: "Return UserID",
			fields: fields{
				email: new(MockEmail),
			},
			want: new(MockEmail),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &user{
				id:    tt.fields.id,
				name:  tt.fields.name,
				email: tt.fields.email,
			}
			got := u.Email()
			assert.Equal(t, tt.want, got, "user.Email() = %v, want %v", got, tt.want)
		})
	}
}

func TestParseUserName(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    UserName
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return UserName",
			args: args{
				value: "testUser",
			},
			want:    userName("testUser"),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when user name is invalid",
			args: args{
				value: "ABCDFGHIJK1234567890ABCDFGHIJK1234567890ABCDFGHIJK1234567890ABCDE",
			},
			want: userName(""),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "invalid user name")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseUserName(tt.args.value)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_userName_Validate(t *testing.T) {
	tests := []struct {
		name    string
		n       userName
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Return UserName when user name length is 1 (Min)",
			n:       userName("t"),
			wantErr: assert.NoError,
		},
		{
			name:    "Return UserName when user name length is 64 (Max)",
			n:       userName("ABCDFGHIJK1234567890ABCDFGHIJK1234567890ABCDFGHIJK1234567890ABCD"),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when user name is too short",
			n:    userName(""),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "invalid user name")
			},
		},
		{
			name: "Return error when user name is too long",
			n:    userName("ABCDFGHIJK1234567890ABCDFGHIJK1234567890ABCDFGHIJK1234567890ABCDE"),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				return assert.EqualError(tt, e, "invalid user name")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.n.Validate()
			tt.wantErr(t, err)
		})
	}
}

func Test_userName_String(t *testing.T) {
	tests := []struct {
		name string
		n    userName
		want string
	}{
		{
			name: "Return user name as string",
			n:    userName("testuser"),
			want: "testuser",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.n.String()
			assert.Equal(t, tt.want, got)
		})
	}
}
