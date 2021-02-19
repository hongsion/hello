package hello_test

import (
	"testing"

	"github.com/hongsion/hello"
)

func TestHello(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "test", want: "Hello World!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hello.Hello(); got != tt.want {
				t.Errorf("Hello() = %v, want %v", got, tt.want)
			}
		})
	}
}
