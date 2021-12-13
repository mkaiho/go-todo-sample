package rdb

import "github.com/mkaiho/go-todo-sample/entity"

/** ID **/
type IDGenerator interface {
	GeneratID() (entity.ID, error)
}
