package media

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func FindMediaFiles(rootDir string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(rootDir, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		if entry.IsDir() {
			return nil
		}

		if IsMediaFile(path) {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("scan folder %q for media files: %w", rootDir, err)
	}

	sort.Strings(files)
	return files, nil
}
