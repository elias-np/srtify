package assets

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
)

type RuntimePaths struct {
	FFmpeg  string
	FFprobe string
	Whisper string
	Model   string
}

type requiredRuntimeFile struct {
	Label      string
	Candidates []string
}

var requiredRuntimeFiles = []requiredRuntimeFile{
	{Label: "ffmpeg", Candidates: []string{"ffmpeg", "ffmpeg.exe"}},
	{Label: "ffprobe", Candidates: []string{"ffprobe", "ffprobe.exe"}},
	{Label: "whisper", Candidates: []string{"whisper", "whisper.exe"}},
	{Label: "model", Candidates: []string{"ggml-base.bin"}},
}

func ExtractRuntimeFiles(baseDir string) (RuntimePaths, error) {
	files := RuntimeFiles()
	paths, err := runtimePaths(baseDir, files)
	if err != nil {
		return RuntimePaths{}, err
	}

	for _, file := range files {
		if err := writeRuntimeFile(baseDir, file); err != nil {
			return RuntimePaths{}, err
		}
	}

	return paths, nil
}

func writeRuntimeFile(baseDir string, file EmbeddedFile) error {
	targetPath := filepath.Join(baseDir, file.Name)
	err := os.WriteFile(targetPath, file.Bytes, os.FileMode(file.Mode()))
	if err != nil {
		return fmt.Errorf("write embedded file %q to %q: %w", file.Name, targetPath, err)
	}

	return nil
}

func runtimePaths(baseDir string, files []EmbeddedFile) (RuntimePaths, error) {
	nameToPath := map[string]string{}
	for _, file := range files {
		nameToPath[file.Name] = filepath.Join(baseDir, file.Name)
	}

	resolved := map[string]string{}
	missing := make([]string, 0, len(requiredRuntimeFiles))
	for _, required := range requiredRuntimeFiles {
		resolvedName, ok := firstResolvedName(nameToPath, required.Candidates)
		if !ok {
			missing = append(missing, required.Label)
			continue
		}

		resolved[required.Label] = resolvedName
	}

	if len(missing) > 0 {
		return RuntimePaths{}, fmt.Errorf(
			"embedded runtime files incomplete for %q/%q: missing %v; available=%v",
			runtime.GOOS,
			runtime.GOARCH,
			missing,
			availableRuntimeNames(files),
		)
	}

	return RuntimePaths{
		FFmpeg:  resolved["ffmpeg"],
		FFprobe: resolved["ffprobe"],
		Whisper: resolved["whisper"],
		Model:   resolved["model"],
	}, nil

}

func firstResolvedName(pathByName map[string]string, names []string) (string, bool) {
	for _, name := range names {
		resolved, ok := pathByName[name]
		if ok {
			return resolved, true
		}
	}

	return "", false
}

func availableRuntimeNames(files []EmbeddedFile) []string {
	names := make([]string, 0, len(files))
	for _, file := range files {
		names = append(names, file.Name)
	}

	sort.Strings(names)
	return names
}
