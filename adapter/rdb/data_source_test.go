package rdb

import (
	"testing"

	"github.com/mkaiho/go-todo-sample/usecase"
	"github.com/stretchr/testify/assert"
)

func Test_query_SelectFieldNames(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "Return column names",
			fields: fields{
				columnNames: []string{"id", "name"},
			},
			want: []string{"id", "name"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := q.SelectFieldNames()
			assert.Equal(t, tt.want, got, "query.SelectFieldNames() = %v, want %v", got, tt.want)
		})
	}
}

func Test_query_TableName(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return table name",
			fields: fields{
				tableName: "dummy_table",
			},
			want: "dummy_table",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := q.TableName()
			assert.Equal(t, tt.want, got, "query.TableName() = %v, want %v", got, tt.want)
		})
	}
}

func Test_query_Join(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return join query",
			fields: fields{
				joinQuery: "JOIN dummy t2 on t1.id = t2.id",
			},
			want: "JOIN dummy t2 on t1.id = t2.id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := q.Join()
			assert.Equal(t, tt.want, got, "query.Join() = %v, want %v", got, tt.want)
		})
	}
}

func Test_query_JoinArgs(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "Return join query args",
			fields: fields{
				joinQueryArgs: []interface{}{"dummy", 999},
			},
			want: []interface{}{"dummy", 999},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := q.JoinArgs()
			assert.Equal(t, tt.want, got, "query.JoinArgs() = %v, want %v", got, tt.want)
		})
	}
}

func Test_query_Where(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return where query",
			fields: fields{
				whereQuery: "id = ?",
			},
			want: "id = ?",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := q.Where()
			assert.Equal(t, tt.want, got, "query.Where() = %v, want %v", got, tt.want)
		})
	}
}

func Test_query_WhereArgs(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		{
			name: "Return where query args",
			fields: fields{
				whereQueryArgs: []interface{}{"dummy", 999},
			},
			want: []interface{}{"dummy", 999},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &query{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := q.WhereArgs()
			assert.Equal(t, tt.want, got, "query.WhereArgs() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewQueryBuilder(t *testing.T) {
	tests := []struct {
		name string
		want QueryBuilder
	}{
		{
			name: "Return QueryBuilder",
			want: &queryBuilder{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewQueryBuilder()
			assert.Equal(t, tt.want, got, "NewQueryBuilder() = %v, want %v", got, tt.want)
		})
	}
}

func Test_queryBuilder_Select(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	type args struct {
		columnNames []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		{
			name: "Return QueryBuilder set column names",
			fields: fields{
				columnNames: make([]string, 0),
			},
			args: args{
				columnNames: []string{"id", "name"},
			},
			want: &queryBuilder{
				columnNames:    []string{"id", "name"},
				tableName:      "",
				joinQuery:      "",
				joinQueryArgs:  make([]interface{}, 0),
				whereQuery:     "",
				whereQueryArgs: make([]interface{}, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := &queryBuilder{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := qb.Select(tt.args.columnNames...)
			assert.Equal(t, tt.want, got, "queryBuilder.Select() = %v, want %v", got, tt.want)
		})
	}
}

func Test_queryBuilder_From(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	type args struct {
		tableName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		{
			name: "Return QueryBuilder set table name",
			fields: fields{
				columnNames: make([]string, 0),
			},
			args: args{
				tableName: "dummy_table",
			},
			want: &queryBuilder{
				columnNames:    make([]string, 0),
				tableName:      "dummy_table",
				joinQuery:      "",
				joinQueryArgs:  make([]interface{}, 0),
				whereQuery:     "",
				whereQueryArgs: make([]interface{}, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := &queryBuilder{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := qb.From(tt.args.tableName)
			assert.Equal(t, tt.want, got, "queryBuilder.From() = %v, want %v", got, tt.want)
		})
	}
}

func Test_queryBuilder_Join(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		{
			name: "Return QueryBuilder set join query and args",
			fields: fields{
				columnNames: make([]string, 0),
			},
			args: args{
				query: "INNER JOIN dummy_table t1 on t1.id = ?",
				args:  []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
			want: &queryBuilder{
				columnNames:    make([]string, 0),
				tableName:      "",
				joinQuery:      "INNER JOIN dummy_table t1 on t1.id = ?",
				joinQueryArgs:  []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
				whereQuery:     "",
				whereQueryArgs: make([]interface{}, 0),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := &queryBuilder{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := qb.Join(tt.args.query, tt.args.args...)
			assert.Equal(t, tt.want, got, "queryBuilder.Join() = %v, want %v", got, tt.want)
		})
	}
}

func Test_queryBuilder_Where(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	type args struct {
		query string
		args  []interface{}
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   QueryBuilder
	}{
		{
			name: "Return QueryBuilder set where query and args",
			fields: fields{
				columnNames: make([]string, 0),
			},
			args: args{
				query: "id = ?",
				args:  []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
			want: &queryBuilder{
				columnNames:    make([]string, 0),
				tableName:      "",
				joinQuery:      "",
				joinQueryArgs:  make([]interface{}, 0),
				whereQuery:     "id = ?",
				whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := &queryBuilder{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got := qb.Where(tt.args.query, tt.args.args...)
			assert.Equal(t, tt.want, got, "queryBuilder.Where() = %v, want %v", got, tt.want)
		})
	}
}

func Test_queryBuilder_Build(t *testing.T) {
	type fields struct {
		columnNames    []string
		tableName      string
		joinQuery      string
		joinQueryArgs  []interface{}
		whereQuery     string
		whereQueryArgs []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		want    Query
		wantErr func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool
	}{
		{
			name: "Return Query",
			fields: fields{
				columnNames:    []string{"id", "name"},
				tableName:      "dummy_table t1",
				joinQuery:      "INNER JOIN dummy_table t2 ON t2.code = ?",
				joinQueryArgs:  []interface{}{"A001"},
				whereQuery:     "t1.id = ?",
				whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
			want: &query{
				columnNames:    []string{"id", "name"},
				tableName:      "dummy_table t1",
				joinQuery:      "INNER JOIN dummy_table t2 ON t2.code = ?",
				joinQueryArgs:  []interface{}{"A001"},
				whereQuery:     "t1.id = ?",
				whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				return assert.Nil(tt, e)
			},
		},
		{
			name: "Return error when column names is not set",
			fields: fields{
				columnNames:    []string{},
				tableName:      "dummy_table t1",
				joinQuery:      "INNER JOIN dummy_table t2 ON t2.code = ?",
				joinQueryArgs:  []interface{}{"A001"},
				whereQuery:     "t1.id = ?",
				whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(tt, wantCode, e.Code(), "queryBuilder.Build() got.Code() = %v, want.Code() %v", e.Code(), wantCode) {
					return false
				}
				wantError := "unknown error is occured: select columns are required"
				return assert.Equal(tt, wantError, e.Error(), "queryBuilder.Build() got.Error() = %v, want.Error() %v", e.Error(), wantError)
			},
		},
		{
			name: "Return error when table names is not set",
			fields: fields{
				columnNames:    []string{"id", "name"},
				tableName:      "",
				joinQuery:      "INNER JOIN dummy_table t2 ON t2.code = ?",
				joinQueryArgs:  []interface{}{"A001"},
				whereQuery:     "t1.id = ?",
				whereQueryArgs: []interface{}{"a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"},
			},
			want: nil,
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				wantCode := usecase.ErrCodeUnkown
				if !assert.Equal(tt, wantCode, e.Code(), "queryBuilder.Build() got.Code() = %v, want.Code() %v", e.Code(), wantCode) {
					return false
				}
				wantError := "unknown error is occured: table name is required"
				return assert.Equal(tt, wantError, e.Error(), "queryBuilder.Build() got.Error() = %v, want.Error() %v", e.Error(), wantError)
			},
		},
		{
			name: "Return Query when optional fields are not set",
			fields: fields{
				columnNames:    []string{"id", "name"},
				tableName:      "dummy_table t1",
				joinQuery:      "",
				joinQueryArgs:  []interface{}{},
				whereQuery:     "",
				whereQueryArgs: []interface{}{},
			},
			want: &query{
				columnNames:    []string{"id", "name"},
				tableName:      "dummy_table t1",
				joinQuery:      "",
				joinQueryArgs:  []interface{}{},
				whereQuery:     "",
				whereQueryArgs: []interface{}{},
			},
			wantErr: func(tt assert.TestingT, e usecase.UseCaseError, i ...interface{}) bool {
				return assert.Nil(tt, e)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			qb := &queryBuilder{
				columnNames:    tt.fields.columnNames,
				tableName:      tt.fields.tableName,
				joinQuery:      tt.fields.joinQuery,
				joinQueryArgs:  tt.fields.joinQueryArgs,
				whereQuery:     tt.fields.whereQuery,
				whereQueryArgs: tt.fields.whereQueryArgs,
			}
			got, err := qb.Build()
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got, "queryBuilder.Build() got1 = %v, want %v", got, tt.want)
		})
	}
}
