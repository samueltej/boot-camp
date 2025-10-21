package main

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	tmpDir := t.TempDir()

	oldWd, _ := os.Getwd()
	t.Cleanup(func() { os.Chdir(oldWd) })
	os.Chdir(tmpDir)

	mdFile := "testfile.md"
	mdContent := []byte("# Test\nThis is a markdown test.")
	err := os.WriteFile(mdFile, mdContent, 0644)
	if err != nil {
		t.Fatalf("failed to create markdown file: %v", err)
	}
	
	err = run(mdFile, "testfile")
	if err != nil {
		t.Fatalf("run() returned an unexpected error: %v", err)
	}
	outputPath := filepath.Join(tmpDir, "md", "testfile.html")

	content, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("failed to read output file: %v", err)
	}

	if !bytes.HasPrefix(content, []byte(header)) {
		t.Error("output file missing header")
	}

	if !bytes.HasSuffix(content, []byte(footer)) {
		t.Error("output file missing footer")
	}

}
func TestParseContent(t *testing.T) {
	mdPath := filepath.Join("testdata", "test.md")
	goldenPath := filepath.Join("testdata", "test_golden.html")

	mdContent, err := os.ReadFile(mdPath)
	if err != nil {
		t.Fatalf("cannot read markdown file: %v", err)
	}

	got := parseContent(mdContent)
	if len(got) == 0 {
		t.Fatal("parseContent returned an empty slice")
	}

	want, err := os.ReadFile(goldenPath)
	if err != nil {
		t.Fatalf("cannot read golden file: %v", err)
	}

	if !bytes.Equal(got, want) {
		t.Errorf("output does not match golden file.\n\nExpected:\n%s\n\nGot:\n%s", string(want), string(got))
	}
}
