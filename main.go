package main

import (
	"fmt"
	"os"
	"path"
	"strconv"

	"github.com/dtbead/telegram-chan/config"
	"github.com/dtbead/telegram-chan/file"
	"github.com/dtbead/telegram-chan/telegram"
)

type Config struct {
	ChannelID      string
	APIToken       string
	WatchPath      string
	EncodeToPath   string
	UploadFromPath string
	EncodeSpeed    string
}

const DefaultConfigPath = "config"

var Bot *telegram.Bot

func main() {
	var c config.Config
	if len(os.Args) > 1 {
		if os.Args[1] != "" {
			c = config.LoadConfig(os.Args[1])
		}
	} else {
		c = config.LoadConfig(DefaultConfigPath)
	}

	cs := Config{
		ChannelID:      c.Get("CHANNELID"),
		APIToken:       c.Get("APITOKEN"),
		WatchPath:      c.Get("WATCHPATH"),
		EncodeToPath:   c.Get("ENCODEPATH"),
		UploadFromPath: c.Get("UPLOADPATH"),
		EncodeSpeed:    c.Get("ENCODESPEED"),
	}

	Bot, err := telegram.NewBot(cs.APIToken)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	StartVideoWatch(Bot, cs)
}

func StartVideoWatch(Bot *telegram.Bot, cs Config) {
	if err := file.WatchFolder([]string{cs.EncodeToPath, cs.WatchPath, cs.UploadFromPath}, func(s string) { videoWatch(Bot, cs, s) }); err != nil {
		fmt.Printf("StartVideoWatch/file.WatchFolder err: %v\n", err)
	}
}

func videoWatch(Bot *telegram.Bot, cs Config, filepath string) {
	if file.IsDirectory(filepath) {
		fmt.Println("ignoring directory")
		return
	}

	if !file.IsVideo(filepath) {
		fmt.Println("ignoring non-video")
		return
	}

	dir := path.Dir(filepath)
	if dir == cs.UploadFromPath && file.GetSize(filepath) > 512 {
		channelID, _ := strconv.Atoi(cs.ChannelID)

		rsp, err := Bot.SendVideo(filepath, int64(channelID))
		if err != nil {
			fmt.Printf("telegram error: %v\n", err.Error())
		}

		fmt.Printf("uploaded video to %s with id %d\n", rsp.Chat.Title, rsp.MessageID)
	}

	if dir == cs.WatchPath {
		if err := file.Encode(filepath, cs.EncodeToPath, cs.EncodeSpeed); err != nil {
			fmt.Printf("main.videoWatch()/file.Encode() err: %v\n", err.Error())
		}
	}
}
