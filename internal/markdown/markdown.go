package markdown

import (
	"fmt"
	"os"
	"strings"

	"github.com/russross/blackfriday/v2"
	"gopkg.in/yaml.v2"
)

type Metadata struct {
	Layout   string `yaml:"layout"`
	PageType string `yaml:"page-type"`
	Title    string `yaml:"title"`
}

type MarkdownFile struct {
	Path     string
	Metadata Metadata
	Content  []byte
}

func NewMarkdownFile(inputPath string) (*MarkdownFile, error) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to read file %s: %v", inputPath, err)
	}

	m := &MarkdownFile{
		Path: inputPath,
	}

	if err := m.parse(input); err != nil {
		return nil, fmt.Errorf("Failed to parse markdown file %s: %v", inputPath, err)
	}

	return m, nil
}

func (m MarkdownFile) GetTitle() string {
	return m.Metadata.Title
}

func (m MarkdownFile) HTML() []byte {
	return convertMarkdownToHTML(m.Content)
}

func convertMarkdownToHTML(input []byte) []byte {
	return blackfriday.Run(input)
}

func (m *MarkdownFile) parse(data []byte) error {
	parts := strings.SplitN(string(data), "---", 3)
	if len(parts) < 3 {
		return fmt.Errorf("invalid markdown file format")
	}

	var metadata Metadata
	if err := yaml.Unmarshal([]byte(parts[1]), &metadata); err != nil {
		return fmt.Errorf("failed to parse metadata: %v", err)
	}

	m.Metadata = metadata
	m.Content = []byte(parts[2])

	return nil
}
