package core

import (
	"strings"
)

type markdownRemover struct {
	// text *string
	// pos  int64
	r *strings.Reader
}

func (o *markdownRemover) Read(p []byte) (int, error) {
	// @TODO: implement
	return o.r.Read(p)
}

func (o *markdownRemover) ReadRune() (rune, int, error) {
	// @TODO: implement
	return o.r.ReadRune()
}

func (o *markdownRemover) UnreadRune() error {
	// @TODO: implement
	return o.r.UnreadRune()
}

func NewMarkdownRemover(t string) *markdownRemover {
	return &markdownRemover{r: strings.NewReader(t)}
}
