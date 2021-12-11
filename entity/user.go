package entity

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	userNameMinLength = 1
	userNameMaxLength = 64
)

/** UserID**/
type UserID interface {
	IsEmpty() bool
	Validate() error
	String() string
}

/** Name **/
type UserName interface {
	Validate() error
	String() string
}

func ParseUserName(value string) (UserName, error) {
	name := userName(value)
	if err := name.Validate(); err != nil {
		return userName(""), err
	}
	return name, nil
}

type userName string

func (n userName) Validate() error {
	if length := len(n); length < userNameMinLength || length > userNameMaxLength {
		return errors.New("invalid user name")
	}
	return nil
}

func (n userName) String() string {
	return string(n)
}

/** Email **/
type Email interface {
	Validate() error
	String() string
}

func ParseEmail(value string) (email, error) {
	parsed := email(value)
	if err := parsed.Validate(); err != nil {
		return "", err
	}
	return parsed, nil
}

type email string

func (em email) String() string {
	return string(em)
}
func (em email) Validate() error {
	const pattern = "[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\\.[a-zA-Z0-9-]+)*$"
	if len(em) == 0 {
		return errors.New("email is empty")
	}
	if regex := regexp.MustCompile(pattern); !regex.Match([]byte(em.String())) {
		return fmt.Errorf("email format is invalid: \"%v\"", em)
	}
	return nil
}

/** UserID**/

type User interface {
	UserID() UserID
	Name() UserName
	Email() Email
}
type user struct {
	id    UserID
	name  UserName
	email Email
}

func NewUser(
	id UserID,
	name UserName,
	email Email,
) (User, error) {
	if err := id.Validate(); !id.IsEmpty() && err != nil {
		return nil, err
	}
	if err := name.Validate(); err != nil {
		return nil, err
	}
	if err := email.Validate(); err != nil {
		return nil, err
	}
	return &user{
		id:    id,
		name:  name,
		email: email,
	}, nil
}

func (u *user) UserID() UserID {
	return u.id
}

func (u *user) Name() UserName {
	return u.name
}

func (u *user) Email() Email {
	return u.email
}
