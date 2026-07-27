//go:build linux && arm64

package assets

import "embed"

//go:embed files/linux/arm64/*
var linuxArm64Files embed.FS

func RuntimeFiles() []EmbeddedFile {
	return runtimeFilesFromDirectory(
		linuxArm64Files,
		"files/linux/arm64",
		executableSet("ffmpeg", "ffprobe", "whisper"),
	)
}
