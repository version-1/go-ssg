package template

import "path/filepath"
import "github.com/yourproject/markdown"

func GetTemplateFilePath(projectRoot string, metadata markdown.Metadata) string {
	return filepath.Join(projectRoot, "templates", metadata.Layout+".tmpl.html")
}
