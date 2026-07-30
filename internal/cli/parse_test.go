package cli

import (
	"strings"
	"testing"
)

func TestParseReturnsDefaultSRTFormat(t *testing.T) {
	options, err := Parse([]string{"video.mp4"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if options.Format != FormatSRT {
		t.Fatalf("expected format %q, got %q", FormatSRT, options.Format)
	}
}

func TestParseAcceptsTextOutput(t *testing.T) {
	options, err := Parse([]string{"-txt", "video.mp4"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if options.Format != FormatTXT {
		t.Fatalf("expected format %q, got %q", FormatTXT, options.Format)
	}
}

func TestParseAcceptsTextOutputWhenInputComesFirst(t *testing.T) {
	options, err := Parse([]string{"video.mp4", "-txt"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if options.Format != FormatTXT {
		t.Fatalf("expected format %q, got %q", FormatTXT, options.Format)
	}

	if options.Input != "video.mp4" {
		t.Fatalf("expected input %q, got %q", "video.mp4", options.Input)
	}
}

func TestParseAcceptsMixedOrderWithLanguageAndOutput(t *testing.T) {
	options, err := Parse([]string{"video.mp4", "-txt", "--language", "pt", "--output", "saida/final"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if options.Language != "pt" {
		t.Fatalf("expected language %q, got %q", "pt", options.Language)
	}

	if options.OutputBase != "saida/final" {
		t.Fatalf("expected output base %q, got %q", "saida/final", options.OutputBase)
	}
}

func TestParseEnablesDebugMode(t *testing.T) {
	options, err := Parse([]string{"video.mp4", "--debug"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if !options.Debug {
		t.Fatal("expected debug flag to be true")
	}
}

func TestParseRequiresInput(t *testing.T) {
	_, err := Parse([]string{})
	if err == nil {
		t.Fatal("expected missing input error, got nil")
	}

	if !strings.Contains(err.Error(), "expected input media path") {
		t.Fatalf("expected missing input error to mention media path, got %q", err.Error())
	}
}

func TestParseAllowsHelpWithoutInput(t *testing.T) {
	options, err := Parse([]string{"-h"})
	if err != nil {
		t.Fatalf("expected help parse to succeed, got error %v", err)
	}

	if !options.ShowHelp {
		t.Fatal("expected help flag to be true")
	}
}

func TestParseAllowsRecursiveWithoutInput(t *testing.T) {
	options, err := Parse([]string{"-r"})
	if err != nil {
		t.Fatalf("expected recursive parse to succeed, got error %v", err)
	}

	if !options.Recursive {
		t.Fatal("expected recursive flag to be true")
	}

	if options.Input != "" {
		t.Fatalf("expected empty input to default to current directory later, got %q", options.Input)
	}
}

func TestParseAcceptsRecursiveWithSourceAndOutputFolders(t *testing.T) {
	options, err := Parse([]string{"-r", "--output", "saida", "midias"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if options.Input != "midias" {
		t.Fatalf("expected input folder %q, got %q", "midias", options.Input)
	}

	if options.OutputBase != "saida" {
		t.Fatalf("expected output folder %q, got %q", "saida", options.OutputBase)
	}
}

func TestParseAcceptsGranularityFlag(t *testing.T) {
	options, err := Parse([]string{"video.mp4", "-g", "5"})
	if err != nil {
		t.Fatalf("expected parse to succeed, got error %v", err)
	}

	if options.Granularity != 5 {
		t.Fatalf("expected granularity 5, got %d", options.Granularity)
	}
}

func TestParseRejectsGranularityOutOfRange(t *testing.T) {
	_, err := Parse([]string{"video.mp4", "-g", "6"})
	if err == nil {
		t.Fatal("expected out-of-range granularity error, got nil")
	}

	if !strings.Contains(err.Error(), "expected granularity between 0 and 5") {
		t.Fatalf("expected granularity range error, got %q", err.Error())
	}
}
