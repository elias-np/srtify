package cli

const Version = "0.1.0"

func HelpText() string {
	return `srtify - generate subtitles from video files

Usage:
  srtify [flags] <video-file>

Flags:
  -srt                    generate .srt subtitle file (default)
  -txt                    generate plain .txt transcript
  -l, --language          transcription language (pt, en, auto) [default: auto]
  -o, --output            base name or full path for output files
  -f, --force-transcribe  ignore existing subtitles and transcribe audio
  --debug                 show debug logs during processing
  --verbose               alias for --debug
  -v, --version           show version
  -h, --help              show this help

Behavior:
  - input file path: relative to current directory by default
  - output file path: saved in current directory by default
  - you can pass full paths for input and output when needed
`
}
