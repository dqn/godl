package godl

import (
	"os"
	"testing"
)

func TestDownload(t *testing.T) {
	dest := "test.json"
	err := Download("https://api.github.com/users/dqn", dest, false)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(dest)
}

func TestDownloadShowProgressWithDenominator(t *testing.T) {
	dest := "test.zip"
	err := Download("https://www.post.japanpost.jp/zipcode/dl/kogaki/zip/ken_all.zip", dest, true)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(dest)
}

func TestDownloadShowProgressWithoutDenominator(t *testing.T) {
	dest := "test.json"
	err := Download("https://github.com/dqn", dest, true)
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(dest)
}
