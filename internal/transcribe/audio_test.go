package transcribe

import (
	"reflect"
	"testing"

	"srtify/internal/cli"
)

func TestWhisperArgsUseAutoLanguageWhenConfigured(t *testing.T) {
	request := Request{
		ModelPath:  "model.bin",
		OutputBase: "out/base",
		Language:   "auto",
		Format:     cli.FormatTXT,
	}

	got := whisperArgs(request, "input.wav")
	want := []string{"-m", "model.bin", "-f", "input.wav", "-of", "out/base", "-otxt", "-l", "auto"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected args %v, got %v", want, got)
	}
}

func TestWhisperArgsUseExplicitLanguage(t *testing.T) {
	request := Request{
		ModelPath:  "model.bin",
		OutputBase: "out/base",
		Language:   "pt",
		Format:     cli.FormatSRT,
	}

	got := whisperArgs(request, "input.wav")
	want := []string{"-m", "model.bin", "-f", "input.wav", "-of", "out/base", "-osrt", "-l", "pt"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected args %v, got %v", want, got)
	}
}
