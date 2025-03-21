package markdown

import (
	"fmt"
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

func ParseMarkdownFile(data []byte) (*MarkdownFile, error) {
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
