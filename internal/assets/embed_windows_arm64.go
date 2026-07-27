//go:build windows && arm64

package assets

import "embed"

//go:embed files/windows/arm64/*
var windowsArm64Files embed.FS

func RuntimeFiles() []EmbeddedFile {
	return runtimeFilesFromDirectory(
		windowsArm64Files,
		"files/windows/arm64",
		executableSet("ffmpeg.exe", "ffprobe.exe", "whisper.exe"),
	)
}
