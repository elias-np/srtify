package media

import (
	"os"
	"path/filepath"
	"strings"
)

var commonAudioExtensions = map[string]bool{
	".wav":  true,
	".mp3":  true,
	".flac": true,
	".ogg":  true,
	".oga":  true,
	".m4a":  true,
	".aac":  true,
	".opus": true,
	".wma":  true,
	".aiff": true,
	".aif":  true,
	".alac": true,
	".amr":  true,
	".mka":  true,
	".caf":  true,
}

var commonVideoExtensions = map[string]bool{
	".mp4":  true,
	".mkv":  true,
	".avi":  true,
	".mov":  true,
	".webm": true,
	".flv":  true,
	".wmv":  true,
	".m4v":  true,
	".mpg":  true,
	".mpeg": true,
	".3gp":  true,
	".ts":   true,
}

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

func IsLikelyAudioInput(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return commonAudioExtensions[extension]
}

func IsMediaFile(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return commonAudioExtensions[extension] || commonVideoExtensions[extension]
}

func withoutExtension(path string) string {
	name := filepath.Base(path)
	extension := filepath.Ext(name)
	return strings.TrimSuffix(name, extension)
}
