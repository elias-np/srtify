package subtitle

import (
	"fmt"
	"os/exec"
)

func ExtractFirst(ffmpegPath string, inputPath string, outputPath string) error {
	args := extractArgs(inputPath, outputPath)
	command := exec.Command(ffmpegPath, args...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("extract subtitle from %q to %q failed: %w: %s", inputPath, outputPath, err, output)
	}

	return nil
}

func extractArgs(inputPath string, outputPath string) []string {
	return []string{
		"-y",
		"-i", inputPath,
		"-map", "0:s:0",
		outputPath,
	}
}