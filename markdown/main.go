package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const header = `<!DOCTYPE html>
  <html>
    <head>
      <meta http-equiv="content-type" content="text/html; charset=utf-8" />
      <title>Markdown Preview Tool</title>
    </head>
    <body>
`

const footer = `
    </body>
  </html>
`

func main() {
	in := flag.String("in", "", "path to the markdown file to parse")
	out := flag.String("out", "", "output HTML file name (without extension)")
	flag.Parse()

	if err := run(*in, *out, os.Stdout); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(in, out string, writer io.Writer) error {
	if in == "" {
		return fmt.Errorf("the -in flag is required")
	}

	inputData, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	content := parseContent(inputData)

	var outputName string
	if out != "" {
		outputName = "md/" + out + ".html"
	} else {
		outputName = "md/md*.html"
	}

	if err := os.MkdirAll(filepath.Dir(outputName), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}

	if out == "" {
		tempFile, err := os.CreateTemp("md", "md*.html")
		if err != nil {
			return fmt.Errorf("failed to create temporary file: %v", err)
		}
		outputName = tempFile.Name()
		tempFile.Close()
	}

	if err := saveHTML(outputName, content); err != nil {
		return fmt.Errorf("failed to save HTML file: %v", err)
	}

	fmt.Fprintln(writer, outputName)
	return nil
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	htmlContent := append([]byte(header), body...)
	htmlContent = append(htmlContent, []byte(footer)...)

	return htmlContent
}

func saveHTML(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("could not write to file %s: %v", filename, err)
	}
	return nil
}