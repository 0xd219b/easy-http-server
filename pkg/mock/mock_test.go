package mock

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestGenHandler(t *testing.T) {
	type args struct {
		url    string
		method string
		resp   string
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
		{
			name: "test GET",
			args: args{
				url:    "/test",
				method: "GET",
				resp:   `{"name":"test"}`,
			},
			want: func(c *gin.Context) {
				c.JSON(http.StatusOK, `{"name":"test"}`)
			},
		},
		{
			name: "test POST",
			args: args{
				url:    "/test",
				method: "POST",
				resp:   `{"name":"test"}`,
			},
			want: func(c *gin.Context) {
				var f interface{}
				err := json.Unmarshal([]byte(`{"name":"test"}`), &f)
				if err != nil {
					panic(err)
				}
				c.JSON(http.StatusCreated, f)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenHandler(tt.args.url, tt.args.method, tt.args.resp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenHandler() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
