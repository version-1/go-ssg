package content

import (
	"github.com/yourproject/markdown"
	"github.com/yourproject/template"
	"github.com/yourproject/fileutils"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/russross/blackfriday/v2"
	"github.com/microcosm-cc/bluemonday"
)


func sanitizeHTML(input []byte) []byte {
	policy := bluemonday.UGCPolicy()
	return policy.SanitizeBytes(input)
}

func ProcessMarkdownFile(projectRoot, inputPath, outputDir string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", inputPath, err)
	}

	markdownFile, err := markdown.ParseMarkdownFile(input)
	if err != nil {
		log.Fatalf("Failed to parse markdown file %s: %v", inputPath, err)
	}

	templateFilePath := template.GetTemplateFilePath(projectRoot, markdownFile.Metadata)
	templateContent, err := os.ReadFile(templateFilePath)
	if err != nil {
		log.Fatalf("Failed to read template file %s: %v", templateFilePath, err)
	}

	output := markdown.ConvertMarkdownToHTML(markdownFile.Content)

	output = sanitizeHTML(output)
	// TODO: Implement stylesheet and javascript replacement logic
	finalContent := strings.ReplaceAll(string(templateContent), "{{ args.content }}", string(output))
	finalContent = strings.ReplaceAll(finalContent, "{{ args.stylesheet }}", "")
	finalContent = strings.ReplaceAll(finalContent, "{{ args.javascript }}", "")
	finalContent = strings.ReplaceAll(finalContent, "{{ args.title }}", markdownFile.Metadata.Title)

	outputFilePath := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(inputPath), ".md")+".html")
	if err := fileutils.WriteHTMLToFile(outputFilePath, []byte(finalContent)); err != nil {
		log.Fatalf("Failed to write file %s: %v", outputFilePath, err)
	}
}
