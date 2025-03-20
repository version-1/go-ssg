package content

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/russross/blackfriday/v2"
)

type Metadata struct {
	Layout   string `yaml:"layout"`
	PageType string `yaml:"page-type"`
	Title    string `yaml:"title"`
}

type MarkdownFile struct {
	Metadata Metadata
	Content  []byte
}

func ConvertMarkdownToHTML(input []byte) []byte {
	return blackfriday.Run(input)
}

func GetTemplateFilePath(projectRoot string, metadata Metadata) string {
	return filepath.Join(projectRoot, "templates", metadata.Layout+".tmpl.html")
}

func ProcessMarkdownFile(inputPath, outputDir string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", inputPath, err)
	}

	markdownFile, err := parseMarkdownFile(input)
	if err != nil {
		log.Fatalf("Failed to parse markdown file %s: %v", inputPath, err)
	}

	output := ConvertMarkdownToHTML(markdownFile.Content)

	outputFilePath := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(inputPath), ".md")+".html")
	if err := WriteHTMLToFile(outputFilePath, output); err != nil {
		log.Fatalf("Failed to write file %s: %v", outputFilePath, err)
	}
}
func parseMarkdownFile(data []byte) (*MarkdownFile, error) {
	parts := strings.SplitN(string(data), "---", 3)
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid markdown file format")
	}

	var metadata Metadata
	if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
		return nil, fmt.Errorf("failed to parse metadata: %v", err)
	}

	return &MarkdownFile{
		Metadata: metadata,
		Content:  []byte(parts[2]),
	}, nil
}
func WriteHTMLToFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", filePath, err)
	}
	fmt.Printf("Converted to %s\n", filePath)
	return nil
}
