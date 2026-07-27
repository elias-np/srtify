//go:build windows && amd64

package assets

import "embed"

//go:embed files/windows/amd64/*
var windowsAmd64Files embed.FS

func RuntimeFiles() []EmbeddedFile {
	return runtimeFilesFromDirectory(
		windowsAmd64Files,
		"files/windows/amd64",
		executableSet("ffmpeg.exe", "ffprobe.exe", "whisper.exe"),
	)
}
