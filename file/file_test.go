package file

import (
	"testing"
)

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
		{"very malformed path", args{`E:\\\\\Videos\\\telegram\\\\\\\\enc`}, `E:\Videos\telegram\enc`},
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

func Test_IsVideo(t *testing.T) {
	type args struct {
		ext string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"video extension", args{".mp4"}, true},
		{"video windows path", args{`C:\Users\foo\Videos\Captures\replay_2024.04.24-20.26.mp4`}, true},
		{"non video", args{".txt"}, false},
		{"garbage input", args{"hkjdsft789345"}, false},
		{"blank", args{""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsVideo(tt.args.ext); got != tt.want {
				t.Errorf("IsVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}
