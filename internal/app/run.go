package app

import (
	"fmt"
	"os"
	"path/filepath"

	"srtify/internal/assets"
	"srtify/internal/cli"
	"srtify/internal/media"
	"srtify/internal/subtitle"
	"srtify/internal/transcribe"
)

func Run(options cli.Options) error {
	if options.ShowHelp {
		fmt.Print(cli.HelpText())
		return nil
	}

	if options.ShowVersion {
		fmt.Println(cli.Version)
		return nil
	}

	outputPath, err := processVideo(options)
	if err != nil {
		return err
	}

	fmt.Printf("finalizado com sucesso: arquivo gerado em %q\n", outputPath)
	return nil
}

func processVideo(options cli.Options) (string, error) {
	debugLog(options, "iniciando processamento do arquivo %q", options.Input)
	baseDir, err := os.MkdirTemp("", "srtify-*")
	if err != nil {
		return "", fmt.Errorf("create runtime directory with pattern %q: %w", "srtify-*", err)
	}
	defer os.RemoveAll(baseDir)
	debugLog(options, "diretorio temporario criado em %q", baseDir)

	paths, err := assets.ExtractRuntimeFiles(baseDir)
	if err != nil {
		return "", err
	}
	debugLog(options, "assets extraidos com sucesso")

	outputPath, err := generateOutput(options, paths)
	if err != nil {
		return "", err
	}

	if err := ensureOutputWasGenerated(outputPath); err != nil {
		return "", err
	}

	debugLog(options, "saida confirmada em %q", outputPath)
	return outputPath, nil
}

func generateOutput(options cli.Options, paths assets.RuntimePaths) (string, error) {
	base := media.OutputBase(options.Input, options.OutputBase)
	outputBase, err := media.ResolveOutputBase(base)
	if err != nil {
		return "", fmt.Errorf("resolve output base %q: %w", base, err)
	}

	outputPath := expectedOutputPath(options, outputBase)
	debugLog(options, "arquivo de saida esperado em %q", outputPath)

	if err := ensureOutputDirectory(outputPath); err != nil {
		return "", err
	}

	if shouldExtractSubtitle(options) {
		debugLog(options, "modo extracao de legenda embutida")
		return outputPath, subtitle.ExtractFirst(paths.FFmpeg, options.Input, outputPath)
	}

	if options.Format == cli.FormatSRT && media.IsLikelyAudioInput(options.Input) {
		debugLog(options, "entrada de audio detectada; usando transcricao para gerar srt")
	}

	debugLog(options, "modo transcricao por whisper")
	request := transcribe.RequestFromOptions(options, paths, outputBase)
	return outputPath, transcribe.Audio(request)
}

func shouldExtractSubtitle(options cli.Options) bool {
	if options.ForceTranscribe || options.Format != cli.FormatSRT {
		return false
	}

	return !media.IsLikelyAudioInput(options.Input)
}

func expectedOutputPath(options cli.Options, outputBase string) string {
	if options.Format == cli.FormatTXT {
		return outputBase + ".txt"
	}

	return outputBase + ".srt"
}

func ensureOutputWasGenerated(outputPath string) error {
	if _, err := os.Stat(outputPath); err != nil {
		return fmt.Errorf("process finished but output file %q was not found: %w", outputPath, err)
	}

	return nil
}

func ensureOutputDirectory(outputPath string) error {
	directory := filepath.Dir(outputPath)
	if directory == "." || directory == "" {
		return nil
	}

	if err := os.MkdirAll(directory, 0o755); err != nil {
		return fmt.Errorf("create output directory %q: %w", directory, err)
	}

	return nil
}

func debugLog(options cli.Options, format string, values ...any) {
	if !options.Debug && !options.Verbose {
		return
	}

	fmt.Fprintf(os.Stderr, "[debug] "+format+"\n", values...)
}
