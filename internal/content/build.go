package content

import (
	"fmt"

	"github.com/version-1/go-ssg/internal/config"
	"github.com/version-1/go-ssg/internal/fileutils"
	"github.com/version-1/go-ssg/internal/markdown"
	"github.com/version-1/go-ssg/internal/template"
)

func Build(c *config.Config, inputPath string) error {
	m, err := markdown.NewMarkdownFile(inputPath)
	if err != nil {
		return err
	}

	tmpl, err := template.NewTemplate(c, m.Metadata.Layout)
	if err != nil {
		return err
	}
	fmt.Printf("tmpl: %s\n", tmpl.Stylesheet)

	pageContent := NewPage(tmpl, m.Metadata, m.HTML()).Render()

	outputFilePath := c.GetOutputPagePath(inputPath, ".md")
	if err := fileutils.WriteFile(outputFilePath, []byte(pageContent)); err != nil {
		return fmt.Errorf("Failed to write file %s: %v", outputFilePath, err)
	}

	return nil
}
