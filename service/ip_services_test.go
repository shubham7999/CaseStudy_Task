package service

import (
	"os"
	"testing"
)

func TestHandleRequest(t *testing.T) {
   
    tests := []struct {
        name string
    }{
        {
           
            name: "test-1",
        },
    }
    for _, tt := range tests {

        t.Run(tt.name, func(t *testing.T) {
            os.Setenv("local", "1")
			HandleRequests()
            os.Unsetenv("local")

        })

    }
}