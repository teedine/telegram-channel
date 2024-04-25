package file

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	ffmpeg_go "github.com/u2takey/ffmpeg-go"
)

func Encode(filename, output, speed string, quality int) error {
	scaleFilterArgs := ffmpeg_go.Args{
		"min(1280,iw):min(720,ih)",
		"force_original_aspect_ratio=decrease",
		"flags=lanczos",
	}

	unSharpFilterArgs := ffmpeg_go.Args{
		"chroma_amount=1.2",
		"luma_amount=0",
	}

	HEVCargs := ffmpeg_go.KwArgs{
		"flags":    "+global_header",
		"movflags": "faststart",
		"preset":   speed,
		"pix_fmt":  "yuv420p",
		"c:v":      "libx265",
		"crf":      quality,
		"r":        60,
		"c:a":      "aac",
		"b:a":      "64k",
		"map":      "0:a",
	}

	stream := ffmpeg_go.Input(filename).Filter("scale", scaleFilterArgs).Filter("unsharp", unSharpFilterArgs)

	h := stream.Hash()
	str := intToHashString(h)
	fmt.Printf("encode.encode hash:%s", str)

	if !strings.HasSuffix(output, "/") || !strings.HasSuffix(output, "\\") {
		output = output + "/"
	}
	output = output + str + ".mkv"

	err := stream.OverWriteOutput().Output(output, HEVCargs).ErrorToStdOut().Run()
	if err != nil {
		return err
	}

	return nil
}

func GetDuration(filename string) int {
	out, _ := ffmpeg_go.Probe(filename)
	f := unmarshalJson(out)

	if value, ok := f["format"]["duration"].(string); ok {
		i, _ := strconv.ParseFloat(value, 64)
		return int(i) // we don't care that much about precision
	} else {
		return -1
	}
}

func unmarshalJson(j string) map[string]map[string]interface{} {
	var fields map[string]map[string]interface{}
	err := json.Unmarshal([]byte(j), &fields)
	if err != nil {
		fmt.Printf("encode.unmarshalJson err: %v\n", err)
	}

	return fields
}

func intToHashString(i int) string {
	var b []byte
	b = binary.AppendVarint(b, int64(i))

	return byteToString(b)
}

func byteToString(h []byte) string {
	return hex.EncodeToString(h)
}
