package core

import (
	"io"
	"strings"
)

type markdownRemover struct {
	// text *string
	// pos  int64
	r *strings.Reader
}

func (r *markdownRemover) Read(p []byte) (int, error) {
	// @TODO: implement
	return r.Read(p)
}

func NewMarkdownRemover(t string) io.Reader {
	return &markdownRemover{r: strings.NewReader(t)}
}
