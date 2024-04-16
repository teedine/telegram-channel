package file

import "testing"

func Test_cleanPath(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"malformed path", args{`E:\\Videos\\telegram\\enc`}, `E:\Videos\telegram\enc`},
		{"clean path", args{`E:\Videos\telegram\enc`}, `E:\Videos\telegram\enc`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := cleanPath(tt.args.s); got != tt.want {
				t.Errorf("cleanPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
