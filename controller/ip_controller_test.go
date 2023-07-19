package controller

import (
	"ipProject/config"
	"ipProject/model"
	"net/http"
	"os"
	"testing"
)

type testResponseWriter struct{}

func (t *testResponseWriter) Header() http.Header {
	return make(map[string][]string)
}
func (t *testResponseWriter) Write([]byte) (int, error) {
	return 9, nil
}
func (t *testResponseWriter) WriteHeader(statusCode int) {
	
}
func TestHandleIpCount(t *testing.T) {
    os.Setenv("local", "1")
    type args struct {
        w http.ResponseWriter
        r *http.Request
    }
    
    tests := []struct {
        name string
        args args
        mock  func()
    }{
        {
           
            name: "test-1",
            args: args{
                w: &testResponseWriter{},
                r: &http.Request{},
            },
            mock: func ()  {
                os.Setenv("threshold", "1")
            },
        },
        {
           
            name: "test-2",
            args: args{
                w: &testResponseWriter{},
                r: &http.Request{},
            },
            mock: func ()  {
               config.InputJson = model.IpConfig{

                  Data : []model.IpData{

                    {
                       Ip :"121.22.1.2",
                       HostName: "ksks",
                       Active: true,
                    },
                    {
                        Ip :"121.22.1.2",
                        HostName: "hsh",
                        Active: false,
                    },
                },
               }
            },
        },
    }
    
    for _, tt := range tests {

        tt.mock()
        t.Run(tt.name, func(t *testing.T) {
            HandleIpCount(tt.args.w, tt.args.r)
        })
        os.Unsetenv("threshold")

    }

    os.Unsetenv("local")

}

func TestHandlePing(t *testing.T) {
    os.Setenv("local", "1")
    type args struct {
        w http.ResponseWriter
        r *http.Request
    }
    
    tests := []struct {
        name string
        args args

    }{
        {
           
            name: "test-1",
            args: args{
                w: &testResponseWriter{},
                r: &http.Request{},
            },

        },
        {
           
            name: "test-2",
            args: args{
                w: &testResponseWriter{},
                r: &http.Request{},
            },

        },
    }
    
    for _, tt := range tests {

        t.Run(tt.name, func(t *testing.T) {
            HandlePing(tt.args.w, tt.args.r)
        })
        os.Unsetenv("threshold")

    }

    os.Unsetenv("local")

}