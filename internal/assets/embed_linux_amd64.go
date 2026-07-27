//go:build linux && amd64

package assets

import "embed"

//go:embed files/linux/amd64/*
var linuxAmd64Files embed.FS

func RuntimeFiles() []EmbeddedFile {
	return runtimeFilesFromDirectory(
		linuxAmd64Files,
		"files/linux/amd64",
		executableSet("ffmpeg", "ffprobe", "whisper"),
	)
}
