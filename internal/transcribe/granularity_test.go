package transcribe

import (
	"reflect"
	"testing"
)

func TestGranularityArgsReturnsNilForDefaultLevel(t *testing.T) {
	got := granularityArgs(0)
	if got != nil {
		t.Fatalf("expected no extra args for default granularity, got %v", got)
	}
}

func TestGranularityArgsUsesMaxLenOneForWordByWordLevel(t *testing.T) {
	got := granularityArgs(5)
	want := []string{"-ml", "1", "-sow"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected args %v, got %v", want, got)
	}
}

func TestGranularityArgsUsesSmallGroupLenForLevelFour(t *testing.T) {
	got := granularityArgs(4)
	want := []string{"-ml", "20", "-sow"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected args %v, got %v", want, got)
	}
}

func TestWhisperArgsAppendGranularityFlags(t *testing.T) {
	request := Request{
		ModelPath:   "model.bin",
		OutputBase:  "out/base",
		Language:    "auto",
		Format:      "srt",
		Granularity: 5,
	}

	got := whisperArgs(request, "input.wav")
	want := []string{"-m", "model.bin", "-f", "input.wav", "-of", "out/base", "-osrt", "-l", "auto", "-ml", "1", "-sow"}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected args %v, got %v", want, got)
	}
}
