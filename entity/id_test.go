package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseID(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    ID
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "Return ID when value is valid",
			args: args{
				value: "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c",
			},
			want:    ID("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when value format is invalid",
			args: args{
				value: "dummy_id",
			},
			want: ID(""),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "id format is invalid: \"dummy_id\""
				return assert.EqualError(t, e, want, "ParseID() error = %v, wantErr %v", want, e)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseID(tt.args.value)
			if !tt.wantErr(t, err, "ParseID() error = %v", err) {
				return
			}
			assert.Equal(t, tt.want, got, "ParseID() = %v, want %v", got, tt.want)
		})
	}
}

func TestID_IsEmpty(t *testing.T) {
	tests := []struct {
		name string
		id   ID
		want bool
	}{
		{
			name: "Return true when id is empty",
			id:   ID(""),
			want: true,
		},
		{
			name: "Return false when id is not empty",
			id:   ID("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"),
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.id.IsEmpty()
			assert.Equal(t, tt.want, got, "ID.IsEmpty() = %v, want %v", got, tt.want)
		})
	}
}

func TestID_Validate(t *testing.T) {
	tests := []struct {
		name    string
		id      ID
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "Return no error when value is valid and not empty",
			id:      ID("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"),
			wantErr: assert.NoError,
		},
		{
			name:    "Return no error when value is empty",
			id:      ID(""),
			wantErr: assert.NoError,
		},
		{
			name: "Return error when value format is invalid",
			id:   ID("dummy_id"),
			wantErr: func(tt assert.TestingT, e error, i ...interface{}) bool {
				want := "id format is invalid: \"dummy_id\""
				return assert.EqualError(tt, e, want, "ID.Validate() error = %v, wantErr %v", e, want)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.id.Validate()
			tt.wantErr(t, err)
		})
	}
}

func TestID_String(t *testing.T) {
	tests := []struct {
		name string
		id   ID
		want string
	}{
		{
			name: "Return string",
			id:   ID("a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c"),
			want: "a24e61a9-6ff2-c3bf-4dc8-02ee1491ce1c",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.id.String()
			assert.Equal(t, tt.want, got, "ID.String() = %v, want %v", got, tt.want)
		})
	}
}
