package dictionary_test

import (
	"errors"
	"io"
	"strings"

	dict "github.com/ClarksonCJ/radius/dictionary"
)

type MemoryFile struct {
	Filename string
	Contents string

	r io.Reader
}

func (m *MemoryFile) Read(p []byte) (n int, err error) {
	if m.r == nil {
		m.r = strings.NewReader(m.Contents)
	}
	return m.r.Read(p)
}

func (m *MemoryFile) Close() error {
	return nil
}

func (m *MemoryFile) Name() string {
	return m.Filename
}

type MemoryOpener []MemoryFile

func (m MemoryOpener) OpenFile(name string) (dict.File, error) {
	for _, file := range m {
		if file.Filename == name {
			return &file, nil
		}
	}
	return nil, errors.New("unknown file " + name)
}

var files = MemoryOpener{
	{
		Filename: "simple.dict",
		Contents: `
ATTRIBUTE User-Name 1 string
ATTRIBUTE User-Password 2 octets encrypt=1

ATTRIBUTE Mode 127 integer
VALUE Mode Full 1
VALUE Mode Half 2

ATTRIBUTE ARAP-Challenge-Response 84 octets[8]
`,
	},

	{
		Filename: "recursive_1.dict",
		Contents: `
$INCLUDE recursive_2.dict
`,
	},
	{
		Filename: "recursive_2.dict",
		Contents: `
$INCLUDE recursive_1.dict
`,
	},
}
