package content

import (
	"strings"

	// Add this import

	"github.com/microcosm-cc/bluemonday"
	"github.com/version-1/go-ssg/internal/markdown"
	"github.com/version-1/go-ssg/internal/template"
)

type Page struct {
	Template *template.Template
	Metadata markdown.Metadata
	Content  []byte
}

func (c Page) getTemplate() string {
	return string(c.Template.Content)
}

func (c Page) getContent() string {
	return string(sanitizeHTML(c.Content))
}

func (c Page) getStylesheet() string {
	return string(sanitizeStylesheet(c.Template.Stylesheet))
}

func (c Page) getJavascript() string {
	return string(c.Template.Javascript)
}

func NewPage(tmpl *template.Template, metadata markdown.Metadata, content []byte) *Page {
	return &Page{
		Template: tmpl,
		Metadata: metadata,
		Content:  content,
	}
}

func (c Page) Render() []byte {
	finalContent := strings.ReplaceAll(c.getTemplate(), "{{ args.content }}", c.getContent())
	finalContent = strings.ReplaceAll(finalContent, "{{ args.stylesheet }}", c.getStylesheet())
	finalContent = strings.ReplaceAll(finalContent, "{{ args.javascript }}", c.getJavascript())
	finalContent = strings.ReplaceAll(finalContent, "{{ args.title }}", c.Metadata.Title)

	return []byte(finalContent)
}

func sanitizeHTML(input []byte) []byte {
	policy := bluemonday.UGCPolicy()
	return policy.SanitizeBytes(input)
}

func sanitizeStylesheet(input []byte) []byte {
	return input
}
