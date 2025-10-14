package main

import (
	"flag"
	"fmt"
	"os"
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
	out := flag.String("out", "", "output HTML file name (without extension)")
	flag.Parse()

	if err := run(*out); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(out string) error {
	if out == "" {
		return fmt.Errorf("the -out flag is required")
	}

	content := []byte(header + footer)
	filename := "md/" + out + ".html"

	if err := saveHTML(filename, content); err != nil {
		return fmt.Errorf("failed to save HTML file: %v", err)
	}

	fmt.Printf("File successfully created: %s\n", filename)
	return nil
}

func saveHTML(filename string, data []byte) error {
	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("could not write to file %s: %v", filename, err)
	}
	return nil
}
