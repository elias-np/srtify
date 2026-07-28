package media

import (
	"path/filepath"
	"testing"
)

func TestOutputPathUsesInputNameWhenOutputBaseIsEmpty(t *testing.T) {
	got := OutputPath("sample.video.mp4", "", "srt")
	want := "sample.video.srt"

	if got != want {
		t.Fatalf("expected output path %q, got %q", want, got)
	}
}

func TestOutputPathUsesExplicitOutputBase(t *testing.T) {
	got := OutputPath("sample.mp4", "legendas/final", "txt")
	want := "legendas/final.txt"

	if got != want {
		t.Fatalf("expected output path %q, got %q", want, got)
	}
}

func TestResolveOutputBaseUsesCurrentDirectoryForRelativePath(t *testing.T) {
	resolved, err := ResolveOutputBase("resultado/final")
	if err != nil {
		t.Fatalf("expected resolve output base to succeed, got error %v", err)
	}

	if !filepath.IsAbs(resolved) {
		t.Fatalf("expected resolved output to be absolute, got %q", resolved)
	}

	wantSuffix := filepath.Join("resultado", "final")
	if filepath.Clean(resolved[len(resolved)-len(wantSuffix):]) != filepath.Clean(wantSuffix) {
		t.Fatalf("expected resolved output to end with %q, got %q", wantSuffix, resolved)
	}
}

func TestResolveOutputBaseKeepsAbsolutePath(t *testing.T) {
	abs := filepath.Join(t.TempDir(), "destino", "saida")
	resolved, err := ResolveOutputBase(abs)
	if err != nil {
		t.Fatalf("expected resolve output base to succeed, got error %v", err)
	}

	if filepath.Clean(resolved) != filepath.Clean(abs) {
		t.Fatalf("expected resolved output %q, got %q", abs, resolved)
	}
}

func TestIsLikelyAudioInputReturnsTrueForCommonAudioExtensions(t *testing.T) {
	inputs := []string{
		"podcast.wav",
		"voice.MP3",
		"take.flac",
		"song.m4a",
		"clip.opus",
		"archive.mka",
	}

	for _, input := range inputs {
		if !IsLikelyAudioInput(input) {
			t.Fatalf("expected %q to be detected as audio input", input)
		}
	}
}

func TestIsLikelyAudioInputReturnsFalseForNonAudioExtensions(t *testing.T) {
	inputs := []string{
		"video.mp4",
		"movie.mkv",
		"subtitle.srt",
		"file",
	}

	for _, input := range inputs {
		if IsLikelyAudioInput(input) {
			t.Fatalf("expected %q to not be detected as audio input", input)
		}
	}
}