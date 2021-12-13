package entity

import (
	"fmt"
	"regexp"
)

/** ID **/
type ID string

func ParseID(value string) (ID, error) {
	id := ID(value)
	if err := id.Validate(); err != nil {
		return "", err
	}
	return id, nil
}

func (id ID) IsEmpty() bool {
	return len(id) == 0
}

func (id ID) Validate() error {
	pattern := regexp.MustCompile("^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$")
	if !id.IsEmpty() && !pattern.Match([]byte(id.String())) {
		return fmt.Errorf("id format is invalid: \"%s\"", id.String())
	}
	return nil
}

func (id ID) String() string {
	return string(id)
}
