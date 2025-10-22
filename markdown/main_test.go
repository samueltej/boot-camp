package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunUsingOutFlag(t *testing.T) {
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

	var outputBuffer bytes.Buffer
	err = run(mdFile, "testoutput", &outputBuffer)
	if err != nil {
		t.Fatalf("run() returned an unexpected error: %v", err)
	}

	expectedOutput := "md/testoutput.html"
	if !strings.Contains(outputBuffer.String(), expectedOutput) {
		t.Errorf("expected output to contain %s, got %s", expectedOutput, outputBuffer.String())
	}

	outputPath := filepath.Join(tmpDir, expectedOutput)
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

func TestRunWithoutOutFlag(t *testing.T) {
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

	var outputBuffer bytes.Buffer
	err = run(mdFile, "", &outputBuffer)
	if err != nil {
		t.Fatalf("run() returned an unexpected error: %v", err)
	}

	outputFileName := strings.TrimSpace(outputBuffer.String())
	if outputFileName == "" {
		t.Fatal("no output file name was printed")
	}

	if !strings.HasPrefix(filepath.Base(outputFileName), "md") || !strings.HasSuffix(outputFileName, ".html") {
		t.Errorf("output file name doesn't match expected pattern: %s", outputFileName)
	}

	content, err := os.ReadFile(outputFileName)
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
