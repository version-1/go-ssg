package config

import (
	"path/filepath"
	"strings"
)

type Output struct {
	Dir string
}

func (o Output) PagePath(filePath, ext string) string {
	return filepath.Join(o.Dir, strings.TrimSuffix(filepath.Base(filePath), ext)+".html")
}
