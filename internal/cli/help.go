package cli

const Version = "0.2.0"

func HelpText() string {
	return `srtify - generate subtitles from media files

Usage:
  srtify [flags] <input-media-file>
  srtify -r [flags] [source-folder]

Flags:
  -srt                    generate .srt subtitle file (default)
  -txt                    generate plain .txt transcript
  -l, --language          transcription language (pt, en, auto) [default: auto]
  -o, --output            base name or full path for output files
                          (destination folder when combined with -r)
  -f, --force-transcribe  ignore embedded subtitles and transcribe media
  -r, --recursive         process every media file inside a folder
  -g, --granularity       subtitle timing granularity, 0-5 [default: 0]
                          0 = whisper's natural sentence-length segments
                          1 = long phrases (~10 words)
                          2 = medium phrases (~7 words)
                          3 = short phrases (~5 words)
                          4 = small groups (~3-5 words)
                          5 = word by word
  --debug                 show debug logs during processing
  --verbose               alias for --debug
  -v, --version           show version
  -h, --help              show this help

Behavior:
  - input file path: relative to current directory by default
  - output file path: saved in current directory by default
  - audio inputs are transcribed by default (srt/txt)
  - you can pass full paths for input and output when needed
  - with -r, the source folder defaults to the current directory and
    the destination folder defaults to the source folder; the folder
    structure of the source is mirrored into the destination
  - granularity only affects transcription (whisper); it has no effect
    when srtify extracts an already embedded subtitle track
`
}
