package cli

const Version = "0.1.0"

func HelpText() string {
	return `srtify - generate subtitles from media files

Usage:
  srtify [flags] <input-media-file>

Flags:
  -srt                    generate .srt subtitle file (default)
  -txt                    generate plain .txt transcript
  -l, --language          transcription language (pt, en, auto) [default: auto]
  -o, --output            base name or full path for output files
  -f, --force-transcribe  ignore embedded subtitles and transcribe media
  --debug                 show debug logs during processing
  --verbose               alias for --debug
  -v, --version           show version
  -h, --help              show this help

Behavior:
  - input file path: relative to current directory by default
  - output file path: saved in current directory by default
  - audio inputs are transcribed by default (srt/txt)
  - you can pass full paths for input and output when needed
`
}
