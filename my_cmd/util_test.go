package main

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_normalize(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "basic",
			args: args{
				s: `
key:
  val`,
			},
			want: "key: val\n",
		}, {
			name: "error",
			args: args{
				s: `
key:
val`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := normalize(tt.args.s)
			assert.Equal(t, tt.want, got, "normalize()")
			if assert.Equal(t, tt.wantErr, (err != nil), "normalize() error") == false {
				log.Println(err)
			}
		})
	}
}
