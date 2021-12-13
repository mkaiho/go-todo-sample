package rdb

import (
	"errors"

	"github.com/mkaiho/go-todo-sample/usecase"
)

/** DataSource **/
type DataSource interface {
	Query(dest interface{}, query Query) usecase.UseCaseError
	Execute(sql string, args ...interface{}) usecase.UseCaseError
}

/** Query **/
var _ Query = (*query)(nil)

type Query interface {
	SelectFieldNames() []string
	TableName() string
	Join() string
	JoinArgs() []interface{}
	Where() string
	WhereArgs() []interface{}
}

type query struct {
	columnNames    []string
	tableName      string
	joinQuery      string
	joinQueryArgs  []interface{}
	whereQuery     string
	whereQueryArgs []interface{}
}

func (q *query) SelectFieldNames() []string {
	return q.columnNames
}

func (q *query) TableName() string {
	return q.tableName
}

func (q *query) Join() string {
	return q.joinQuery
}

func (q *query) JoinArgs() []interface{} {
	return q.joinQueryArgs
}

func (q *query) Where() string {
	return q.whereQuery
}

func (q *query) WhereArgs() []interface{} {
	return q.whereQueryArgs
}

/** QueryBuilder **/
var _ QueryBuilder = (*queryBuilder)(nil)

type QueryBuilder interface {
	Select(columnNames ...string) QueryBuilder
	From(tableName string) QueryBuilder
	Join(query string, args ...interface{}) QueryBuilder
	Where(query string, args ...interface{}) QueryBuilder
	Build() (Query, usecase.UseCaseError)
}

func NewQueryBuilder() QueryBuilder {
	return &queryBuilder{}
}

type queryBuilder struct {
	columnNames    []string
	tableName      string
	joinQuery      string
	joinQueryArgs  []interface{}
	whereQuery     string
	whereQueryArgs []interface{}
}

func (qb *queryBuilder) Select(columnNames ...string) QueryBuilder {
	builder := qb.copy()
	builder.columnNames = columnNames
	return builder
}

func (qb *queryBuilder) From(tableName string) QueryBuilder {
	builder := qb.copy()
	builder.tableName = tableName
	return builder
}

func (qb *queryBuilder) Join(query string, args ...interface{}) QueryBuilder {
	builder := qb.copy()
	builder.joinQuery = query
	builder.joinQueryArgs = args
	return builder
}

func (qb *queryBuilder) Where(query string, args ...interface{}) QueryBuilder {
	builder := qb.copy()
	builder.whereQuery = query
	builder.whereQueryArgs = args
	return builder
}

func (qb *queryBuilder) Build() (Query, usecase.UseCaseError) {
	if len(qb.columnNames) == 0 {
		return nil, usecase.NewErrUnknown(errors.New("select columns are required"))
	}
	if len(qb.tableName) == 0 {
		return nil, usecase.NewErrUnknown(errors.New("table name is required"))
	}
	return &query{
		columnNames:    qb.columnNames,
		tableName:      qb.tableName,
		joinQuery:      qb.joinQuery,
		joinQueryArgs:  qb.joinQueryArgs,
		whereQuery:     qb.whereQuery,
		whereQueryArgs: qb.whereQueryArgs,
	}, nil
}

func (qb *queryBuilder) copy() *queryBuilder {
	queryBuilder := queryBuilder{
		tableName:  qb.tableName,
		joinQuery:  qb.joinQuery,
		whereQuery: qb.whereQuery,
	}
	queryBuilder.columnNames = append([]string{}, qb.columnNames...)
	queryBuilder.joinQueryArgs = append([]interface{}{}, qb.joinQueryArgs...)
	queryBuilder.whereQueryArgs = append([]interface{}{}, qb.whereQueryArgs...)
	return &queryBuilder
}
