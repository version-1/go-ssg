package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/version-1/go-ssg/internal/content"
)

func ensureDirExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-ssg <project-root>")
		os.Exit(1)
	}

	projectRoot := os.Args[1]
	inputDir := filepath.Join(projectRoot, "pages")
	outputDir := filepath.Join(projectRoot, "public")

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relativePath, err := filepath.Rel(inputDir, path)
			if err != nil {

				return err
			}
			outputPath := filepath.Join(outputDir, strings.TrimSuffix(relativePath, ".md")+".html")
			if err := ensureDirExists(filepath.Dir(outputPath)); err != nil {
				return err
			}
			content.ConvertMarkdownToHTML(path, outputDirPath)
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", inputDir, err)
	}
}
