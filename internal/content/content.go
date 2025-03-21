package content

import (
	"strings"

	"github.com/version-1/go-ssg/internal/markdown"
)

func CreateFinalContent(templateContent, output []byte, metadata markdown.Metadata) string {
	finalContent := strings.ReplaceAll(string(templateContent), "{{ args.content }}", string(output))
	finalContent = strings.ReplaceAll(finalContent, "{{ args.stylesheet }}", "")
	finalContent = strings.ReplaceAll(finalContent, "{{ args.javascript }}", "")
	finalContent = strings.ReplaceAll(finalContent, "{{ args.title }}", metadata.Title)
	return finalContent
}
