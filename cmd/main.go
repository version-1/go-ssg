package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/version-1/go-ssg/internal/content"
)

func main() {
	inputDir := "content"
	outputDir := "public"

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			content.ConvertMarkdownToHTML(path, outputDir)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", inputDir, err)
	}
}
