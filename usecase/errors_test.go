package usecase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewErrNotFoundEntity(t *testing.T) {
	type args struct {
		name  string
		param interface{}
	}
	tests := []struct {
		name string
		args args
		want UseCaseError
	}{
		{
			name: "Return Not Found Entity Error",
			args: args{
				name:  "dummy",
				param: "dummy_param",
			},
			want: &useCaseError{
				code:    ErrCodeNotFoundEntity,
				message: "dummy entity is not found. param: dummy_param",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewErrNotFoundEntity(tt.args.name, tt.args.param)
			assert.Equal(t, tt.want, got, "NewErrNotFoundEntity() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewErrInvalidUsecaseInput(t *testing.T) {
	type args struct {
		name   string
		fields map[string]interface{}
	}
	tests := []struct {
		name string
		args args
		want UseCaseError
	}{
		{
			name: "Return Invalid Usecase Input Error",
			args: args{
				name: "dummy",
				fields: map[string]interface{}{
					"id": "dummy_id",
				},
			},
			want: &useCaseError{
				code:    ErrCodeInvalidUsecaseInput,
				message: "dummy input is invalid: map[id:dummy_id]",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewErrInvalidUsecaseInput(tt.args.name, tt.args.fields)
			assert.Equal(t, tt.want, got, "NewErrInvalidUsecaseInput() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewErrUnknown(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want UseCaseError
	}{
		{
			name: "Return Unknown Error",
			args: args{
				err: errors.New("dummy error"),
			},
			want: &useCaseError{
				code:    ErrCodeUnkown,
				message: "unknown error is occured: dummy error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewErrUnknown(tt.args.err)
			assert.Equal(t, tt.want, got, "NewErrUnknown() = %v, want %v", got, tt.want)
		})
	}
}

func Test_useCaseError_Code(t *testing.T) {
	type fields struct {
		code    errorCode
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   errorCode
	}{
		{
			name: "Return error code",
			fields: fields{
				code: 999,
			},
			want: 999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &useCaseError{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			got := e.Code()
			assert.Equal(t, tt.want, got, "useCaseError.Code() = %v, want %v", got, tt.want)
		})
	}
}

func Test_useCaseError_Error(t *testing.T) {
	type fields struct {
		code    errorCode
		message string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Return error content",
			fields: fields{
				message: "dummy error",
			},
			want: "dummy error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &useCaseError{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			got := e.Error()
			assert.Equal(t, tt.want, got, "useCaseError.Error() = %v, want %v", got, tt.want)
		})
	}
}

func Test_useCaseError_IsCodeEqualTo(t *testing.T) {
	type fields struct {
		code    errorCode
		message string
	}
	type args struct {
		code errorCode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Return true when error code is equal",
			fields: fields{
				code: 1,
			},
			args: args{
				code: 1,
			},
			want: true,
		},
		{
			name: "Return fa false hen error code is not equal",
			fields: fields{
				code: 1,
			},
			args: args{
				code: 999,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &useCaseError{
				code:    tt.fields.code,
				message: tt.fields.message,
			}
			got := e.IsCodeEqualTo(tt.args.code)
			assert.Equal(t, tt.want, got, "useCaseError.IsCodeEqualTo() = %v, want %v", got, tt.want)
		})
	}
}

func Test_errorCode_IsEqualCodeInError(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		ec   errorCode
		args args
		want bool
	}{
		{
			name: "Return true when error interface match UseCaseError and code is equal",
			ec:   999,
			args: args{
				err: &useCaseError{
					code: 999,
				},
			},
			want: true,
		},
		{
			name: "Return false when error is not match UseCaseError interface",
			ec:   999,
			args: args{
				err: errors.New("dummy"),
			},
			want: false,
		},
		{
			name: "Return false when error code is not equal",
			ec:   999,
			args: args{
				err: &useCaseError{
					code: 1,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ec.IsEqualCodeInError(tt.args.err); got != tt.want {
				t.Errorf("errorCode.IsEqualCodeInError() = %v, want %v", got, tt.want)
			}
		})
	}
}
