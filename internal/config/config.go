package config

import (
	"path/filepath"
)

type Config struct {
	ProjectRoot string
	InputDir    string
	Output      Output
	TemplateDir string
}

func (c *Config) GetInputDir() string {
	return filepath.Join(c.ProjectRoot, c.InputDir)
}

func (c *Config) GetOutputDir() string {
	return filepath.Join(c.ProjectRoot, c.Output.Dir)
}
func (c *Config) GetOutputPagePath(filePath, ext string) string {
	return filepath.Join(c.ProjectRoot, c.Output.PagePath(filePath, ext))
}

func (c *Config) GetTemplateDir() string {
	return filepath.Join(c.ProjectRoot, c.TemplateDir)
}

func (c Config) GetTemplatePath(layoutName string) string {
	return filepath.Join(c.GetTemplateDir(), layoutName)
}

func NewConfig(projectRoot string) *Config {
	return &Config{
		ProjectRoot: projectRoot,
		InputDir:    "pages",
		Output: Output{
			Dir: "public",
		},
		TemplateDir: "templates",
	}
}
