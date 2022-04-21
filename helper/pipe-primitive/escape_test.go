package pipe_primitive

import (
	"reflect"
	"testing"
)

func Test_replaceReverseShellAddress(t *testing.T) {
	type args struct {
		payload []byte
		host    string
		port    uint
	}
	tests := []struct {
		name         string
		args         args
		wantReplaced []byte
		wantErr      bool
	}{
		{
			name: "65534",
			args: args{
				payload: []byte{0xff, 0xfe, 0xfd, 0xfc, ':', 0x5b, 0x25},
				host:    "127.0.0.1",
				port:    uint(65534),
			},
			wantReplaced: []byte{127, 0, 0, 1, ':', 0xff, 0xfe},
			wantErr:      false,
		},
	//	{
	//		name: "23333",
	//		args: args{
	//			payload: []byte{0xff, 0xfe, 0xfd, 0xfc, ':', 0x5b, 0x25},
	//			host:    "127.0.0.1",
	//			port:    uint(23333),
	//		},
	//		wantReplaced: []byte{127, 0, 0, 1, ':', 0x5b, 0x25},
	//		wantErr:      false,
	//	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotReplaced, err := replaceReverseShellAddress(tt.args.payload, tt.args.host, tt.args.port)
			if (err != nil) != tt.wantErr {
				t.Errorf("replaceReverseShellAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotReplaced, tt.wantReplaced) {
				t.Errorf("replaceReverseShellAddress() gotReplaced = %v, want %v", gotReplaced, tt.wantReplaced)
			}
		})
	}
}