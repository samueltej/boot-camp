# Markdown to HTML Converter 🧩

This program converts Markdown (`.md`) files into HTML files.

It uses two open-source libraries:
- [`blackfriday`](https://github.com/russross/blackfriday): Converts Markdown to HTML.
- [`bluemonday`](https://github.com/microcosm-cc/bluemonday): Sanitizes the generated HTML to ensure safety.

---

## ⚙️ How It Works

The program takes an input Markdown file and generates an HTML file that includes:
- A predefined **header**
- The **converted HTML content**
- A predefined **footer**

You can specify both input and output file names through flags.

---

## 🧭 Usage

Run the program using the following command:

```bash
go run main.go -in <input-file> [-out <output-file>]
