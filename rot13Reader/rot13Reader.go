package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func (r rot13Reader) Rot13(b byte, base byte) byte {
	val := (b - base + 13)%26 + base
	return val
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.reader.Read(b)
	for i, value := range(b) {
		if value <= 'Z' && value >= 'A' {
			b[i] = 	r.Rot13(value, 'A')
		} else if value <= 'z' && value >= 'a' {
			b[i] = r.Rot13(value, 'a')
		}
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
