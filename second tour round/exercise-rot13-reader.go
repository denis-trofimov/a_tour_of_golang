package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r13.r.Read(b)
	if err != nil {
		return
	}
	for i, l := range b {
		switch {
		case l >= 'a' && l <= 'm' || l >= 'A' && l <= 'M':
			b[i] += 13
		case l >= 'n' && l <= 'z' || l >= 'N' && l <= 'Z':
			b[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
