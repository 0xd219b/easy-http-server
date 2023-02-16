package mock

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGenHandler(t *testing.T) {
	type args struct {
		method string
		resp   []byte
		status int
	}
	type check struct {
		status int
		body   []byte
	}
	tests := []struct {
		name string
		args args
		want check
	}{
		{
			name: "post test",
			args: args{
				method: "POST",
				resp:   []byte("{\"name\":\"test\"}"),
				status: http.StatusCreated,
			},
			want: check{
				status: http.StatusCreated,
				body:   []byte("{\"name\":\"test\"}"),
			},
		},
		{
			name: "get test",
			args: args{
				method: "GET",
				resp:   []byte("{\"name\":\"test\"}"),
				status: http.StatusOK,
			},
			want: check{
				status: http.StatusOK,
				body:   []byte("{\"name\":\"test\"}"),
			},
		},
		{
			name: "status not set",
			args: args{
				method: "POST",
				resp:   []byte("{\"name\":\"test\"}"),
			},
			want: check{
				status: http.StatusCreated,
				body:   []byte("{\"name\":\"test\"}"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.ReleaseMode)
			g := gin.New()
			got := GenHandler(tt.args.method, tt.args.resp, tt.args.status)
			g.Handle(tt.args.method, "test", got)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(tt.args.method, "/test", w.Body)
			g.ServeHTTP(w, r)
			if w.Code != tt.want.status {
				t.Errorf("GenHandler() = %v, want %v", w.Code, tt.want.status)
			}
			if w.Body.String() != string(tt.want.body) {
				t.Errorf("GenHandler() = %v, want %v", w.Body.String(), string(tt.want.body))
			}
		})
	}
}
