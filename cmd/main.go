package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/version-1/go-ssg/internal/config"
	"github.com/version-1/go-ssg/internal/content"
	"github.com/version-1/go-ssg/internal/fileutils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go-ssg <project-root>")
		os.Exit(1)
	}

	projectRoot := os.Args[1]
	cfg := config.NewConfig(projectRoot)
	inputDir := cfg.GetInputDir()

	err := filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".md") {
			relativePath, err := filepath.Rel(inputDir, path)
			if err != nil {
				return err
			}

			outputPath := cfg.Output.PagePath(relativePath, ".md")
			outputDirPath := filepath.Dir(outputPath)
			if err := fileutils.EnsureDirExists(outputDirPath); err != nil {
				return err
			}

			if err := content.Build(cfg, path); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error walking the path %q: %v\n", inputDir, err)
	}
}
