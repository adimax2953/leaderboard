package server

import (
	"reflect"
	"testing"
)

func TestNewAESEncoder(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *AESEncoder
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "AES-128",
			args:    args{key: "0000000000000000", value: "hello"},
			want:    &AESEncoder{},
			wantErr: false,
		},
		{
			name:    "AES-192",
			args:    args{key: "000000000000000000000000", value: "hello"},
			want:    &AESEncoder{},
			wantErr: false,
		},
		{
			name:    "AES-256",
			args:    args{key: "00000000000000000000000000000000", value: "hello"},
			want:    &AESEncoder{},
			wantErr: false,
		},
		{
			name:    "not either 16, 24, or 32 bytes",
			args:    args{key: "00000000000000000", value: "hello"},
			want:    &AESEncoder{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			tt.want, err = NewAESEncoder(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAESEncoder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				var testdata string
				var testBin []byte

				testdata, err = tt.want.Encrypt(tt.args.value)
				if (err != nil) != tt.wantErr {
					t.Errorf("NewAESEncoder() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("AESEncoder.Encrypt(%v) = %v.", tt.args.value, testdata)

				if (err != nil) == tt.wantErr {
					testBin, err = tt.want.Decrypt(testdata)
				}
				if (err != nil) != tt.wantErr {
					t.Errorf("NewAESEncoder() error = %v, wantErr %v", err, tt.wantErr)
					return
				}
				t.Logf("AESEncoder.Decrypt(%v) = %v.", tt.args.value, string(testBin))

				if !reflect.DeepEqual("hello", tt.args.value) {
					t.Errorf("NewAESEncoder() = %v, want %v", testdata, tt.args.value)
				}
			}
		})
	}
}
