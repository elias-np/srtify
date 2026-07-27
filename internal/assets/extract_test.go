package assets

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEmbeddedFileModeUsesExecutablePermissions(t *testing.T) {
	file := EmbeddedFile{Executable: true}

	if got := file.Mode(); got != 0o755 {
		t.Fatalf("expected executable mode %o, got %o", 0o755, got)
	}
}

func TestEmbeddedFileModeUsesRegularPermissions(t *testing.T) {
	file := EmbeddedFile{Executable: false}

	if got := file.Mode(); got != 0o644 {
		t.Fatalf("expected regular mode %o, got %o", 0o644, got)
	}
}

func TestExtractRuntimeFilesWritesAllAssets(t *testing.T) {
	baseDir := t.TempDir()

	paths, err := ExtractRuntimeFiles(baseDir)
	if err != nil {
		t.Fatalf("expected runtime files to extract, got error %v", err)
	}

	assertFileExists(t, paths.FFmpeg)
	assertFileExists(t, paths.FFprobe)
	assertFileExists(t, paths.Whisper)
	assertFileExists(t, paths.Model)
}

func assertFileExists(t *testing.T, path string) {
	t.Helper()

	if _, err := os.Stat(filepath.Clean(path)); err != nil {
		t.Fatalf("expected file %q to exist, got error %v", path, err)
	}
}
