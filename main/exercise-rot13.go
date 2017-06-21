package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(p []byte) (n int, err error) {
	n, err = rot.r.Read(p)
	for i := 0; i < len(p); i++ {
		if (p[i] < 'N' && p[i] >= 'A') || (p[i] < 'n' && p[i] >= 'a') {
			p[i] += 13
		} else if (p[i] <= 'Z' && p[i] > 'M') || (p[i] <= 'z' && p[i] > 'm') {
			p[i] -= 13
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
