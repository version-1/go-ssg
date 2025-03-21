package content

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/version-1/go-ssg/internal/content"
	"github.com/version-1/go-ssg/internal/fileutils"
	"github.com/version-1/go-ssg/internal/markdown"
	"github.com/version-1/go-ssg/internal/template"

	"github.com/microcosm-cc/bluemonday"
)


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

	// TODO: Implement stylesheet and javascript replacement logic
	finalContent := content.CreateFinalContent(templateContent, output, markdownFile.Metadata)

	outputFilePath := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(inputPath), ".md")+".html")
	if err := fileutils.WriteHTMLToFile(outputFilePath, []byte(finalContent)); err != nil {
		log.Fatalf("Failed to write file %s: %v", outputFilePath, err)
	}
}
