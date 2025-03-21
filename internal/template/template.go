package template

import (
	"os"
	"path/filepath"

	"github.com/version-1/go-ssg/internal/config"
)

type TemplatePaths struct {
	BasePath string
}

func NewTemplatePaths(basePath string) *TemplatePaths {
	return &TemplatePaths{BasePath: basePath}
}

func (t *TemplatePaths) HTML() string {
	return filepath.Join(t.BasePath, "index.html")
}

func (t *TemplatePaths) CSS() string {
	return filepath.Join(t.BasePath, "index.css")
}

func (t *TemplatePaths) JS() string {
	return filepath.Join(t.BasePath, "index.js")
}

type Template struct {
	Content    []byte
	Stylesheet []byte
	Javascript []byte
	Path       string
}

func NewTemplate(c *config.Config, layoutName string) (*Template, error) {
	templateFilePath := c.GetTemplatePath(layoutName)
	paths := NewTemplatePaths(templateFilePath)

	contents := make([][]byte, 3)
	for i, resource := range []struct {
		path     string
		required bool
	}{
		{
			path:     paths.HTML(),
			required: true,
		},
		{
			path:     paths.JS(),
			required: false,
		},
		{
			path:     paths.CSS(),
			required: false,
		},
	} {
		content, err := readFile(resource.path, resource.required)
		if err != nil {
			return nil, err
		}
		contents[i] = content
	}

	return &Template{
		Content:    contents[0],
		Javascript: contents[1],
		Stylesheet: contents[2],
		Path:       templateFilePath,
	}, nil
}

func readFile(filePath string, required bool) ([]byte, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) && !required {
		return nil, nil
	}

	if os.IsNotExist(err) && required {
		return nil, err
	}

	return os.ReadFile(filePath)
}
