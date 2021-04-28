package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(bytes []byte) (int, error) {
	n, err := rot.r.Read(bytes)

	if err == nil {
		for i := range bytes {
			char := bytes[i]

			if char >= 'A' && char < 'Z' {
				bytes[i] = 'A' + (char-'A'+13)%26
			} else if char >= 'a' && char < 'z' {
				bytes[i] = 'a' + (char-'a'+13)%26
			}
		}
	}

	return n, err
}

func rot13() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
