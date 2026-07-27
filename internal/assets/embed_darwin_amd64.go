//go:build darwin && amd64

package assets

import "embed"

//go:embed files/darwin/amd64/*
var darwinAmd64Files embed.FS

func RuntimeFiles() []EmbeddedFile {
	return runtimeFilesFromDirectory(
		darwinAmd64Files,
		"files/darwin/amd64",
		executableSet("ffmpeg", "ffprobe", "whisper"),
	)
}
