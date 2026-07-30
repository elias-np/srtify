package app

import (
	"errors"
	"path/filepath"
	"testing"

	"srtify/internal/cli"
)

func TestResolveFolderPathsDefaultsSourceAndDestToCurrentDirectory(t *testing.T) {
	options := cli.Options{Recursive: true}

	source, dest, err := resolveFolderPaths(options)
	if err != nil {
		t.Fatalf("expected resolve to succeed, got error %v", err)
	}

	if !filepath.IsAbs(source) || !filepath.IsAbs(dest) {
		t.Fatalf("expected absolute source and dest, got %q and %q", source, dest)
	}

	if source != dest {
		t.Fatalf("expected dest to default to source, got source %q dest %q", source, dest)
	}
}

func TestResolveFolderPathsRejectsNonExistentSource(t *testing.T) {
	options := cli.Options{Recursive: true, Input: "this-folder-does-not-exist"}

	_, _, err := resolveFolderPaths(options)
	if err == nil {
		t.Fatal("expected error for missing source folder, got nil")
	}
}

func TestResolveFolderPathsUsesExplicitSourceAndDest(t *testing.T) {
	sourceDir := t.TempDir()
	destDir := t.TempDir()
	options := cli.Options{Recursive: true, Input: sourceDir, OutputBase: destDir}

	source, dest, err := resolveFolderPaths(options)
	if err != nil {
		t.Fatalf("expected resolve to succeed, got error %v", err)
	}

	if filepath.Clean(source) != filepath.Clean(sourceDir) {
		t.Fatalf("expected source %q, got %q", sourceDir, source)
	}

	if filepath.Clean(dest) != filepath.Clean(destDir) {
		t.Fatalf("expected dest %q, got %q", destDir, dest)
	}
}

func TestWithoutExtensionStripsExtensionButKeepsDirectory(t *testing.T) {
	got := withoutExtension(filepath.Join("aulas", "licao1.mp4"))
	want := filepath.Join("aulas", "licao1")

	if got != want {
		t.Fatalf("expected %q, got %q", want, got)
	}
}

func TestReportBatchResultsReturnsNilWhenAllSucceed(t *testing.T) {
	results := []batchResult{{Path: "a.srt"}, {Path: "b.srt"}}

	if err := reportBatchResults(results); err != nil {
		t.Fatalf("expected no error when all files succeed, got %v", err)
	}
}

func TestReportBatchResultsReturnsErrorWhenSomeFail(t *testing.T) {
	results := []batchResult{{Path: "a.srt"}, {Err: errors.New("boom")}}

	if err := reportBatchResults(results); err == nil {
		t.Fatal("expected error when some files fail, got nil")
	}
}
