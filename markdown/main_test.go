package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)


func TestRun(t *testing.T) {
	tmp := t.TempDir()

	oldWd, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(oldWd) })
	os.Chdir(tmp)

	os.Mkdir("md", os.ModePerm)

	err := run("testfile")
	if err != nil {
		t.Fatalf("run() returned an unexpected error: %v", err)
	}

	path := filepath.Join(tmp, "md", "testfile.html")

	content, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("failed to read created file: %v", err)
	}

	if !bytes.HasPrefix(content, []byte(header)) {
		t.Error("file content is missing header")
	}
	if !bytes.HasSuffix(content, []byte(footer)) {
		t.Error("file content is missing footer")
	}
}
