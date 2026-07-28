package app

import (
	"testing"

	"srtify/internal/cli"
)

func TestShouldExtractSubtitleReturnsFalseForAudioInput(t *testing.T) {
	options := cli.Options{
		Input:  "podcast.mp3",
		Format: cli.FormatSRT,
	}

	if shouldExtractSubtitle(options) {
		t.Fatal("expected audio input to use transcription instead of embedded subtitle extraction")
	}
}

func TestShouldExtractSubtitleReturnsTrueForVideoSRTInput(t *testing.T) {
	options := cli.Options{
		Input:  "movie.mp4",
		Format: cli.FormatSRT,
	}

	if !shouldExtractSubtitle(options) {
		t.Fatal("expected video input in srt mode to extract embedded subtitles")
	}
}

func TestShouldExtractSubtitleReturnsFalseWhenForceTranscribeIsEnabled(t *testing.T) {
	options := cli.Options{
		Input:           "movie.mp4",
		Format:          cli.FormatSRT,
		ForceTranscribe: true,
	}

	if shouldExtractSubtitle(options) {
		t.Fatal("expected force transcribe mode to skip embedded subtitle extraction")
	}
}

func TestShouldExtractSubtitleReturnsFalseForTXTOutput(t *testing.T) {
	options := cli.Options{
		Input:  "movie.mp4",
		Format: cli.FormatTXT,
	}

	if shouldExtractSubtitle(options) {
		t.Fatal("expected txt mode to skip embedded subtitle extraction")
	}
}
