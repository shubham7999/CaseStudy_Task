package route

import (
	"os"
	"testing"
)

func TestMockIpData(t *testing.T) {

	os.Setenv("local", "1")
	tests := []struct {
		name string
	}{
		{

			name: "test-1",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			InitRoutes()
		})

	}
	os.Unsetenv("local")
}
