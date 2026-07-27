//go:build darwin && arm64

package assets

import "embed"

//go:embed files/darwin/arm64/*
var darwinArm64Files embed.FS

func RuntimeFiles() []EmbeddedFile {
	return runtimeFilesFromDirectory(
		darwinArm64Files,
		"files/darwin/arm64",
		executableSet("ffmpeg", "ffprobe", "whisper"),
	)
}
