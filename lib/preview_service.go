package lib

import (
	"bytes"
	_ "embed"
	"fmt"
	"html/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

//go:embed preview.html
var rawTemplate string

var t *template.Template = template.Must(template.New("preview").Parse(rawTemplate))

type previewContext struct {
	Title   string
	Content template.HTML
}

type PreviewService struct {
	renderer goldmark.Markdown
}

func NewPreviewService(allowUnsafe bool) PreviewService {
	var markdown goldmark.Markdown = goldmark.New(
		goldmark.WithExtensions(extension.GFM),
		goldmark.WithParserOptions(parser.WithAutoHeadingID()),
	)

	if allowUnsafe == true {
		markdown.Renderer().AddOptions(html.WithUnsafe())
	}

	return PreviewService{
		renderer: markdown,
	}
}

func (service PreviewService) ForFile(filename string) (string, error) {
	content, err := readFile(filename)
	if err != nil {
		return "", err
	}

	converted := bytes.Buffer{}
	if err := service.renderer.Convert(content, &converted); err != nil {
		return "", err
	}

	html := bytes.Buffer{}
	t.Execute(&html, previewContext{
		Title:   fmt.Sprintf("Preview of \"%v\" | Peek", filename),
		Content: template.HTML(converted.String()),
	})

	return html.String(), nil
}
