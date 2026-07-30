package transcribe

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"srtify/internal/cli"
)

func Audio(request Request) error {
	audioPath, cleanup, err := prepareAudioInput(request)
	if err != nil {
		return err
	}
	defer cleanup()
	debugLog(request, "audio de entrada para whisper: %q", audioPath)

	args := whisperArgs(request, audioPath)
	debugLog(request, "executando whisper com saida base %q", request.OutputBase)
	command := exec.Command(request.WhisperPath, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("transcribe %q to %q failed: %w: %s", request.InputPath, request.OutputBase, err, output)
	}

	return nil
}

func whisperArgs(request Request, audioPath string) []string {
	args := []string{
		"-m", request.ModelPath,
		"-f", audioPath,
		"-of", request.OutputBase,
	}

	args = append(args, formatFlag(request.Format))
	args = append(args, languageArgs(request.Language)...)
	return append(args, granularityArgs(request.Granularity)...)
}

func prepareAudioInput(request Request) (string, func(), error) {
	if isNativeWhisperAudio(request.InputPath) {
		debugLog(request, "entrada %q ja e audio nativo", request.InputPath)
		return request.InputPath, func() {}, nil
	}

	tempFile, err := os.CreateTemp("", "srtify-audio-*.wav")
	if err != nil {
		return "", nil, fmt.Errorf("create temp wav for %q: %w", request.InputPath, err)
	}

	tempPath := tempFile.Name()
	if closeErr := tempFile.Close(); closeErr != nil {
		return "", nil, fmt.Errorf("close temp wav %q: %w", tempPath, closeErr)
	}

	args := []string{
		"-y",
		"-i", request.InputPath,
		"-vn",
		"-ac", "1",
		"-ar", "16000",
		tempPath,
	}

	command := exec.Command(request.FFmpegPath, args...)
	debugLog(request, "convertendo midia para wav temporario %q", tempPath)
	output, runErr := command.CombinedOutput()
	if runErr != nil {
		os.Remove(tempPath)
		return "", nil, fmt.Errorf("convert media %q to wav failed: %w: %s", request.InputPath, runErr, output)
	}

	cleanup := func() {
		os.Remove(tempPath)
	}

	return tempPath, cleanup, nil
}

func isNativeWhisperAudio(path string) bool {
	extension := strings.ToLower(filepath.Ext(path))
	return extension == ".wav" || extension == ".mp3" || extension == ".flac" || extension == ".ogg"
}

func debugLog(request Request, format string, values ...any) {
	if !request.Debug {
		return
	}

	fmt.Fprintf(os.Stderr, "[debug] "+format+"\n", values...)
}

func formatFlag(format cli.OutputFormat) string {
	if format == cli.FormatTXT {
		return "-otxt"
	}

	return "-osrt"
}

func languageArgs(language string) []string {
	if language == "auto" {
		return []string{"-l", "auto"}
	}

	return []string{"-l", language}
}
