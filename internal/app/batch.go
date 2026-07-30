package app

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"srtify/internal/assets"
	"srtify/internal/cli"
	"srtify/internal/media"
)

type batchResult struct {
	Path string
	Err  error
}

func runFolder(options cli.Options) error {
	sourceDir, destDir, err := resolveFolderPaths(options)
	if err != nil {
		return err
	}
	debugLog(options, "modo recursivo: origem %q, destino %q", sourceDir, destDir)

	files, err := media.FindMediaFiles(sourceDir)
	if err != nil {
		return err
	}

	if len(files) == 0 {
		return fmt.Errorf("no media files found in folder %q", sourceDir)
	}

	var results []batchResult
	err = withRuntimeAssets(options, func(paths assets.RuntimePaths) error {
		results = processFolderFiles(options, paths, sourceDir, destDir, files)
		return nil
	})
	if err != nil {
		return err
	}

	return reportBatchResults(results)
}

func resolveFolderPaths(options cli.Options) (string, string, error) {
	sourceDir := options.Input
	if sourceDir == "" {
		sourceDir = "."
	}

	absSource, err := filepath.Abs(sourceDir)
	if err != nil {
		return "", "", fmt.Errorf("resolve source folder %q: %w", sourceDir, err)
	}

	info, err := os.Stat(absSource)
	if err != nil || !info.IsDir() {
		return "", "", fmt.Errorf("expected source folder path, got %q", sourceDir)
	}

	destDir := options.OutputBase
	if destDir == "" {
		destDir = absSource
	}

	absDest, err := filepath.Abs(destDir)
	if err != nil {
		return "", "", fmt.Errorf("resolve destination folder %q: %w", destDir, err)
	}

	return absSource, absDest, nil
}

func processFolderFiles(options cli.Options, paths assets.RuntimePaths, sourceDir string, destDir string, files []string) []batchResult {
	results := make([]batchResult, 0, len(files))
	for _, file := range files {
		outputPath, err := processFolderFile(options, paths, sourceDir, destDir, file)
		results = append(results, batchResult{Path: outputPath, Err: err})

		if err != nil {
			fmt.Fprintf(os.Stderr, "erro ao processar %q: %v\n", file, err)
			continue
		}

		fmt.Printf("gerado: %q\n", outputPath)
	}

	return results
}

func processFolderFile(options cli.Options, paths assets.RuntimePaths, sourceDir string, destDir string, file string) (string, error) {
	relPath, err := filepath.Rel(sourceDir, file)
	if err != nil {
		return "", fmt.Errorf("resolve relative path for %q under %q: %w", file, sourceDir, err)
	}

	fileOptions := options
	fileOptions.Input = file
	fileOptions.OutputBase = filepath.Join(destDir, withoutExtension(relPath))

	outputPath, err := generateOutput(fileOptions, paths)
	if err != nil {
		return "", err
	}

	return outputPath, ensureOutputWasGenerated(outputPath)
}

func withoutExtension(path string) string {
	return strings.TrimSuffix(path, filepath.Ext(path))
}

func reportBatchResults(results []batchResult) error {
	failed := 0
	for _, result := range results {
		if result.Err != nil {
			failed++
		}
	}

	succeeded := len(results) - failed
	fmt.Printf("processamento em lote concluido: %d de %d arquivos gerados com sucesso\n", succeeded, len(results))

	if failed > 0 {
		return fmt.Errorf("%d de %d arquivos falharam durante o processamento em lote", failed, len(results))
	}

	return nil
}
