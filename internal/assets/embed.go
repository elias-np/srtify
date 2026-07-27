package assets

import (
	"io/fs"
	"path"
	"sort"
)

func executableSet(names ...string) map[string]bool {
	executables := make(map[string]bool, len(names))
	for _, name := range names {
		executables[name] = true
	}

	return executables
}

func runtimeFilesFromDirectory(
	embeddedFiles fs.FS,
	directory string,
	executables map[string]bool,
) []EmbeddedFile {
	entries, err := fs.ReadDir(embeddedFiles, directory)
	if err != nil {
		return nil
	}

	runtimeFiles := make([]EmbeddedFile, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		name := entry.Name()
		content, readErr := fs.ReadFile(embeddedFiles, path.Join(directory, name))
		if readErr != nil {
			continue
		}

		runtimeFiles = append(runtimeFiles, EmbeddedFile{
			Name:       name,
			Bytes:      content,
			Executable: executables[name],
		})
	}

	sort.Slice(runtimeFiles, func(left int, right int) bool {
		return runtimeFiles[left].Name < runtimeFiles[right].Name
	})

	return runtimeFiles
}
