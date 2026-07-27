package transcribe

import (
	"srtify/internal/assets"
	"srtify/internal/cli"
)

type Request struct {
	FFmpegPath  string
	WhisperPath string
	ModelPath   string
	InputPath   string
	OutputBase  string
	Language    string
	Format      cli.OutputFormat
	Debug       bool
}

func RequestFromOptions(options cli.Options, paths assets.RuntimePaths, outputBase string) Request {
	return Request{
		FFmpegPath:  paths.FFmpeg,
		WhisperPath: paths.Whisper,
		ModelPath:   paths.Model,
		InputPath:   options.Input,
		OutputBase:  outputBase,
		Language:    options.Language,
		Format:      options.Format,
		Debug:       options.Debug || options.Verbose,
	}
}