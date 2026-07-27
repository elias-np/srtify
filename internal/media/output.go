package media

import (
	"os"
	"path/filepath"
	"strings"
)

func OutputBase(inputPath string, outputBase string) string {
	if outputBase != "" {
		return outputBase
	}

	return withoutExtension(inputPath)
}

func OutputPath(inputPath string, outputBase string, extension string) string {
	return OutputBase(inputPath, outputBase) + "." + extension
}

func ResolveOutputBase(base string) (string, error) {
	if filepath.IsAbs(base) {
		return base, nil
	}

	workingDir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return filepath.Join(workingDir, base), nil
}

func withoutExtension(path string) string {
	name := filepath.Base(path)
	extension := filepath.Ext(name)
	return strings.TrimSuffix(name, extension)
}