package config

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
		{

			name: "test-2",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			MockIpData()
		})

	}

	os.Unsetenv("local")
}

func TestSetThreshold(t *testing.T) {

	os.Setenv("local", "1")

	tests := []struct {
		name string
	}{
		{

			name: "test-1",
		},
		{

			name: "test-2",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			SetThreshold()
		})

	}

	os.Unsetenv("local")

}



func TestInit(t *testing.T) {

	os.Setenv("local", "1")

	tests := []struct {
		name string
	}{
		{

			name: "test-1",
		},
		{

			name: "test-2",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			Init()
		})

	}
	os.Unsetenv("local")
}