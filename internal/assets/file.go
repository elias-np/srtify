package assets

type EmbeddedFile struct {
	Name       string
	Bytes      []byte
	Executable bool
}

func (file EmbeddedFile) Mode() uint32 {
	if file.Executable {
		return 0o755
	}

	return 0o644
}
