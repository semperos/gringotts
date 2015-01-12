package gringotts

import (
	"path/filepath"
	"testing"
)

type FileExpectation struct {
	File     string
	Expected bool
}

func TestDoesFileExist(t *testing.T) {
	cases := []FileExpectation{
		{File: "README.md", Expected: true},
		{File: "foo.txt", Expected: false},
	}
	for _, c := range cases {
		got := DoesFileExist(c.File)
		if got != c.Expected {
			t.Errorf("DoesFileExist: Expected '%v' but got '%v' for file '%s'", c.Expected, got, c.File)
		}
	}
}

type DownloadExpectation struct {
	Url       string
	LocalFile string
	IsErr     bool
}

const (
	outDir = "out"
)

func TestDownloadFile(t *testing.T) {
	cases := []DownloadExpectation{
		{
			Url:       "https://api.github.com/users/go-lang/repos",
			LocalFile: filepath.Join(outDir, "go-lang-repos.json"),
			IsErr:     false,
		},
		{
			Url:       "http://example.com/doesnt-exist",
			LocalFile: filepath.Join(outDir, "not-here.txt"),
			IsErr:     true,
		},
	}
	for _, c := range cases {
		localFile, err := DownloadFile(c.Url, c.LocalFile)
		if err != nil {
			if !c.IsErr {
				t.Errorf("DownloadFile: An unexpected error occurred while downloading '%s' to local file '%s': %v", c.Url, c.LocalFile, err)
			}
		} else {
			if c.IsErr {
				t.Errorf("DownloadFile: Expected an error when downloading '%s' to local file '%s' but no error occurred.\n", c.Url, c.LocalFile)
			} else {
				if !DoesFileExist(localFile) {
					t.Errorf("DownloadFile: Local file '%s' was not created after downloading '%s'.", localFile, c.Url)
				}
			}
		}
	}
}
