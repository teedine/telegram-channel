## Description
    telegram-channel is a small telegram bot for encoding and uploading videos to a channel automatically
## Install
    go install github.com/teedine/telegram-channel@latest
## Usage
    ./telegram-channel [config file]

    if no path is given to a configuration file, telegram-channel will assume a file named  `config` exists in working directory.
## Config
    CHANNELID	# channel ID to send content to	    
    APITOKEN	# API token recieved from telegram
    WATCHPATH	# folder path to watch newly created videos
    ENCODEPATH	# folder path to encode videos found in WATCHPATH
    UPLOADPATH	# folder path to upload videos to your channel
### Example
    CHANNELID=-1002094548386
    APITOKEN=3248743566:NJVWY2DEMZ2TKNRzha3xk4TZM5Tgq2tlGU3
    WATCHPATH=C:/Users/foo/Captures
    ENCODEPATH=C:/Users/foo/Captures/enc
    UPLOADPATH=C:/Users/foo/Captures/enc

## Notes
 - untested on Linux
 - does not respect telegram's 20mb video upload limit (though it uses harsh compression to hopefully keep it below 20mb)
 - nor does it try to keep the file size within embeddable auto-play requirements
