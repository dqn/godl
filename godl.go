package godl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	mb     = 1 << 20
	format = "%.2f"
)

type progress struct {
	numerator   int64
	denominator int64
}

func (p *progress) Write(b []byte) (int, error) {
	n := len(b)
	p.numerator += int64(n)

	fmt.Printf("\r%s", strings.Repeat(" ", 35))
	fmt.Printf("\rdownloading... "+format, float64(p.numerator)/mb)

	if p.denominator > 0 {
		fmt.Printf("/"+format, float64(p.denominator)/mb)
	}

	fmt.Print(" MB")

	return n, nil
}

func Download(rawURL, dest string, showProgress bool) error {
	resp, err := http.Get(rawURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer f.Close()

	if showProgress {
		w := &progress{denominator: resp.ContentLength}
		_, err = io.Copy(f, io.TeeReader(resp.Body, w))
		fmt.Println()
	} else {
		_, err = io.Copy(f, resp.Body)
	}

	if err != nil {
		return err
	}

	return nil
}
