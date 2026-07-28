package cli

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func Parse(args []string) (Options, error) {
	options := DefaultOptions()
	normalizedArgs := reorderArgsForFlexibleOrder(args)
	flags := flag.NewFlagSet("srtify", flag.ContinueOnError)
	registerFormatFlags(flags, &options)
	registerLanguageFlags(flags, &options)
	registerOutputFlags(flags, &options)
	registerModeFlags(flags, &options)

	if err := flags.Parse(normalizedArgs); err != nil {
		return Options{}, err
	}

	return parsedOptions(flags.Args(), options)
}

func registerFormatFlags(flags *flag.FlagSet, options *Options) {
	flags.BoolFunc("srt", "generate .srt subtitle file", setFormat(options, FormatSRT))
	flags.BoolFunc("txt", "generate plain .txt transcript", setFormat(options, FormatTXT))
}

func registerLanguageFlags(flags *flag.FlagSet, options *Options) {
	flags.StringVar(&options.Language, "l", options.Language, "transcription language")
	flags.StringVar(&options.Language, "language", options.Language, "transcription language")
}

func registerOutputFlags(flags *flag.FlagSet, options *Options) {
	flags.StringVar(&options.OutputBase, "o", "", "base name for output files")
	flags.StringVar(&options.OutputBase, "output", "", "base name for output files")
}

func registerModeFlags(flags *flag.FlagSet, options *Options) {
	flags.BoolVar(&options.ForceTranscribe, "f", false, "ignore existing subtitles")
	flags.BoolVar(&options.ForceTranscribe, "force-transcribe", false, "ignore existing subtitles")
	flags.BoolVar(&options.Debug, "debug", false, "show debug logs")
	flags.BoolVar(&options.Verbose, "verbose", false, "show debug logs")
	flags.BoolVar(&options.ShowVersion, "v", false, "show version")
	flags.BoolVar(&options.ShowVersion, "version", false, "show version")
	flags.BoolVar(&options.ShowHelp, "h", false, "show help")
	flags.BoolVar(&options.ShowHelp, "help", false, "show help")
}

func reorderArgsForFlexibleOrder(args []string) []string {
	valueFlags := map[string]bool{
		"-l":         true,
		"--language": true,
		"-o":         true,
		"--output":   true,
	}

	flagArgs := make([]string, 0, len(args))
	positionalArgs := make([]string, 0, len(args))

	for index := 0; index < len(args); index++ {
		arg := args[index]
		if arg == "--" {
			positionalArgs = append(positionalArgs, args[index+1:]...)
			break
		}

		if consumesValue(arg, valueFlags) {
			flagArgs = append(flagArgs, arg)
			if !strings.Contains(arg, "=") && index+1 < len(args) {
				flagArgs = append(flagArgs, args[index+1])
				index++
			}
			continue
		}

		if strings.HasPrefix(arg, "-") {
			flagArgs = append(flagArgs, arg)
			continue
		}

		positionalArgs = append(positionalArgs, arg)
	}

	return append(flagArgs, positionalArgs...)
}

func consumesValue(arg string, valueFlags map[string]bool) bool {
	base := arg
	if strings.Contains(arg, "=") {
		base = strings.SplitN(arg, "=", 2)[0]
	}

	return valueFlags[base]
}

func setFormat(options *Options, format OutputFormat) func(string) error {
	return func(string) error {
		options.Format = format
		return nil
	}
}

func parsedOptions(args []string, options Options) (Options, error) {
	if err := applyPositionalInput(args, &options); err != nil {
		return Options{}, err
	}

	return options, validateOptions(options)
}

func applyPositionalInput(args []string, options *Options) error {
	if len(args) > 1 {
		return fmt.Errorf("expected zero or one input file, got %d: %v", len(args), args)
	}

	if len(args) == 1 {
		options.Input = args[0]
	}

	return nil
}

func validateOptions(options Options) error {
	if options.ShowHelp || options.ShowVersion {
		return nil
	}

	if options.Input == "" {
		return errors.New("expected input media path, got empty value")
	}

	return nil
}
