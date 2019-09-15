package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_foo(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "basic",
			args: args{s: "test"},
			want: "test-gorel: test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := foo(tt.args.s)
			assert.Equal(t, tt.want, got, "foo()")
		})
	}
}
