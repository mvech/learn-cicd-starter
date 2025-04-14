package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	type args struct {
		headerKey string
		keyValue  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "TestHappy",
			args: args{
				headerKey: "Authorization",
				keyValue:  "ApiKey THISISTHEKEY",
			},
			want:    "THISISTHEKEY",
			wantErr: false,
		},
		{
			name: "No key",
			args: args{
				headerKey: "Authorization",
				keyValue:  "ApiKey",
			},
			want:    "",
			wantErr: true,
		},
		{
			name:    "No authorization",
			args:    args{},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header := make(http.Header)
			header.Add(tt.args.headerKey, tt.args.keyValue)
			got, err := GetAPIKey(header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
