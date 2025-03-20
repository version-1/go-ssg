package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/version-1/go-ssg/internal/content"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-ssg <project-root>")
		os.Exit(1)
	}

	projectRoot := os.Args[1]
	inputDir := filepath.Join(projectRoot, "content")
	outputDir := filepath.Join(projectRoot, "public")

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
