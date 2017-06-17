package core

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
	"unicode"
	"unicode/utf8"
  "fmt"
)

const (
	titleLength = 30
)

/*
 *	spaceCompressor
 */
type spaceCompressor struct {
	src io.RuneScanner
}

func (c spaceCompressor) Read(p []byte) (int, error) {
	var err error
	written := 0
	spaceCount := 0
	for {
		var r rune
		var width int
		if r, width, err = c.src.ReadRune(); err != nil {
			break
		} else {
			if unicode.IsSpace(r) {
				spaceCount += 1
				if spaceCount > 1 {
					continue
				}
			} else {
				spaceCount = 0
			}

			left := len(p) - written
			if left < width {
				if err = c.src.UnreadRune(); err != nil {
					log.Print("UnreadRune failed", err)
				}
				break
			}

			written += utf8.EncodeRune(p[written:], r)
		}
	}
	return written, err
}

func NewSpaceCompressor(src io.RuneScanner) io.Reader {
	return &spaceCompressor{src}
}

/* Make special splitter for making titles
 * - stops on first line break (does not output the newline character)
 * - emits uninterrupted sequencies of runes such as unicode.IsDigit(r) or unicode.IsLetter(r) is true
 * - total ammount of runes emited does not exceed `titleLength`
 * - if the first word is longer than `titleLength` returns first `titleLength` runes
 * - only the newline rune is skipped
 * - non-word runes are attached to previous word */
func makeSplitFunc() func([]byte, bool) (int, []byte, error) {
	var runesReturned int = 0
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		// log.Printf("Split: atEOF: %v, data: '%v'", atEOF, string(data))
		if runesReturned >= titleLength {
			// if called after titleLength runes already returned
			return 0, nil, bufio.ErrFinalToken
		}

		var nrune int = 1
		for width, i := 0, 0; i < len(data); i, nrune = i+width, nrune+1 {
			// getR := func(r rune, _ int) rune { return r }
			// log.Printf("at %v (%q), nrune: %v, total: %v", i, getR(utf8.DecodeRune(data[i:])), nrune, runesReturned)
			if runesReturned+nrune > titleLength {
				// current word not finished but sum already exceeds titleLength
				if runesReturned == 0 {
					// the first word is longer than `titleLength`, return first `titleLength` runes
					return 0, data[:i], bufio.ErrFinalToken
				}
				return 0, nil, bufio.ErrFinalToken
			}
			var r rune
			r, width = utf8.DecodeRune(data[i:])
			if r == utf8.RuneError {
				if !atEOF {
					// stopped mid-rune, request more data
					return 0, nil, nil
				}
				// input is not a correct utf8 string
				return 0, nil, errors.New("MakeTitle split func: input string is not valid utf8")
			}
			if r == '\n' {
				// do not include newline character
				runesReturned += (nrune - 1)
				return i + width, data[:i], bufio.ErrFinalToken
			}
			if !unicode.IsDigit(r) && !unicode.IsLetter(r) {
				// include the non-word character, no point in keeping it in input,
				// we have not reached `titleLength` yet so we'll put it into output anyway
				runesReturned += nrune
				return i + width, data[:i+width], nil
			}
		}

		if atEOF {
			// reaching end of input is same as reaching non-word symbol
			runesReturned += nrune
			return len(data), data, bufio.ErrFinalToken
		}

		// request more data
		return 0, nil, nil
	}
}

func MakeTitle(text string) string {
	log.Println("MakeTitle for", text)
	remover := NewMarkdownRemover(text)
	compressor := NewSpaceCompressor(remover)
	scanner := bufio.NewScanner(compressor)
	scanner.Split(makeSplitFunc())
	var resBuf bytes.Buffer
	for scanner.Scan() {
		resBuf.Write(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		log.Println("scanner failed:", err)
		return ""
	}
	result := strings.TrimSpace(resBuf.String())

	log.Println("Title is", result)
	return result
}

/* Write to both stdout and log */
func LoggedPrintf(format string, a ...interface{}) (n int, err error) {
  log.Printf(format, a...)
  return fmt.Printf(format, a...)
}

/* Write to both stdout and log */
func LoggedPrintln(a ...interface{}) (n int, err error) {
  log.Println(a...)
  return fmt.Println(a...)
}
