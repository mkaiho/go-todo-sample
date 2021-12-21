package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Return true when value length is 0",
			args: args{
				value: "",
			},
			want: true,
		},
		{
			name: "Return true when value length is greater than 0",
			args: args{
				value: "A",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmptyString(tt.args.value)
			assert.Equal(t, tt.want, got, "IsEmptyString() = %v, want %v", got, tt.want)
		})
	}
}
