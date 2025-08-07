package src

import (
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	type args struct {
		u *User
	}
	tests := []struct {
		name string
		args args
		want *User
	}{
		{
			name: "success",
			args: args{
				u: &User{Name: "surya", Id: "1", Age: 24},
			},
			want: &User{Name: "surya", Id: "1", Age: 24},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUser(tt.args.u); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
