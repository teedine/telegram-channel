package file

import (
	"os"
	"testing"
)

func TestEncode(t *testing.T) {
	type args struct {
		filename string
		output   string
	}
	tests := []struct {
		name       string
		args       args
		outputFile string
		wantErr    bool
	}{
		{"generic", args{"testdata/FPS_test_1080p60_L4.2.mkv", "testdata/enc/"}, "testdata/enc/c599e0c393dde7b770.mkv", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Encode(tt.args.filename, tt.args.output, "veryfast"); (err != nil) != tt.wantErr {
				t.Errorf("Encode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})

		t.Cleanup(func() {
			if err := os.Remove(tt.outputFile); err != nil {
				t.Fatalf("Encode() cleanup fail! error = %v", err)
			}
		})
	}
}

func TestGetDuration(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"generic", args{"testdata/FPS_test_1080p60_L4.2.mkv"}, 20},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDuration(tt.args.filename); got != tt.want {
				t.Errorf("GetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
