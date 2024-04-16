package config

import (
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Config
	}{
		{"typical", args{"testdata/config"}, Config{
			keyValue{"PATH", "C:/Users/foo/Videos/Captures"},
			keyValue{"CHANNELID", "-23455464567"},
			keyValue{"APITOKEN", "2345423556744564536"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LoadConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfig_Get(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		c    Config
		args args
		want string
	}{
		{"exists", Config{keyValue{"PATH", "test"}}, args{"PATH"}, "test"},
		{"missing", Config{keyValue{"foo", "test"}}, args{"bar"}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Get(tt.args.s); got != tt.want {
				t.Errorf("Config.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseConfig(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want struct{ Key, Value string }
	}{
		{"typical", args{"foo=bar"}, keyValue{"foo", "bar"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseConfig(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
