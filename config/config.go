package config

import (
	"bufio"
	"fmt"
	"os"
)

type Config []keyValue

type keyValue struct {
	Key, Value string
}

func (c Config) Get(s string) string {
	for _, v := range c {
		if v.Key == s {
			return v.Value
		}
	}
	return ""
}

func LoadConfig(path string) Config {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	var c Config

	b := bufio.NewScanner(file)
	for b.Scan() {
		c = append(c, parseConfig(b.Text()))
	}

	return c
}

func parseConfig(s string) struct{ Key, Value string } {
	var key, value []rune
	var hit bool = false

	for _, v := range s {
		if v == '=' && !hit {
			hit = true
		} else {
			if hit {
				value = append(value, v)
			} else {
				key = append(key, v)
			}
		}
	}
	return struct {
		Key   string
		Value string
	}{string(key), string(value)}
}
