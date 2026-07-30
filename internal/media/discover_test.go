package media

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindMediaFilesReturnsSortedMediaFilesRecursively(t *testing.T) {
	root := t.TempDir()
	writeFile(t, filepath.Join(root, "b.mp4"))
	writeFile(t, filepath.Join(root, "a.mp3"))
	writeFile(t, filepath.Join(root, "notes.txt"))
	writeFile(t, filepath.Join(root, "sub", "c.wav"))

	got, err := FindMediaFiles(root)
	if err != nil {
		t.Fatalf("expected scan to succeed, got error %v", err)
	}

	want := []string{
		filepath.Join(root, "a.mp3"),
		filepath.Join(root, "b.mp4"),
		filepath.Join(root, "sub", "c.wav"),
	}

	if len(got) != len(want) {
		t.Fatalf("expected files %v, got %v", want, got)
	}

	for index, path := range want {
		if got[index] != path {
			t.Fatalf("expected files %v, got %v", want, got)
		}
	}
}

func writeFile(t *testing.T, path string) {
	t.Helper()

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("create parent dir for %q: %v", path, err)
	}

	if err := os.WriteFile(path, []byte("x"), 0o644); err != nil {
		t.Fatalf("write file %q: %v", path, err)
	}
}
