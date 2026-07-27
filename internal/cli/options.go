package cli

type OutputFormat string

const (
	FormatSRT OutputFormat = "srt"
	FormatTXT OutputFormat = "txt"
)

type Options struct {
	Input           string
	OutputBase      string
	Language        string
	Format          OutputFormat
	ForceTranscribe bool
	Debug           bool
	Verbose         bool
	ShowVersion     bool
	ShowHelp        bool
}

func DefaultOptions() Options {
	return Options{
		Language: "auto",
		Format:   FormatSRT,
	}
}
