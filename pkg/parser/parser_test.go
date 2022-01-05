package parser

import (
	"reflect"
	"testing"
)

func Test_getParsersFromRawBytes(t *testing.T) {
	type args struct {
		rawBytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []Parser
		wantErr bool
	}{
		{
			name: "test getParsersFromRawBytes",
			args: args{
				rawBytes: []byte(`
- url: /test
  method: GET
  resp: "{\"name\":\"test\"}"
- url: /test
  method: POST
  resp: "{\"name\":\"test\"}"
`),
			},
			want: []Parser{
				{
					URL:    "/test",
					Method: "GET",
					Resp:   `{"name":"test"}`,
				},
				{
					URL:    "/test",
					Method: "POST",
					Resp:   `{"name":"test"}`,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getParsersFromRawBytes(tt.args.rawBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("getParsersFromRawBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getParsersFromRawBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
