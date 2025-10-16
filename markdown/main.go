package main

import (
	"flag"
	"fmt"
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

	if err := run(*in, *out); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(in, out string) error {
	if in == "" {
		return fmt.Errorf("the -in flag is required")
	}

	inputData, err := os.ReadFile(in)
	if err != nil {
		return fmt.Errorf("failed to read input file: %v", err)
	}

	body := parseContent(inputData)

	content := append([]byte(header), body...)
	content = append(content, []byte(footer)...)

	var outputName string
	if out == "" {
		base := filepath.Base(in)
		outputName = "md/" + base + ".html"
	} else {
		outputName = "md/" + out + ".html"
	}

	if err := os.MkdirAll(filepath.Dir(outputName), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %v", err)
	}


	if err := saveHTML(outputName, content); err != nil {
		return fmt.Errorf("failed to save HTML file: %v", err)
	}

	fmt.Printf("HTML file successfully created: %s\n", outputName)
	return nil
}

func parseContent(input []byte) []byte {
	output := blackfriday.Run(input)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)
	return body
}

func saveHTML(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("could not write to file %s: %v", filename, err)
	}
	return nil
}
