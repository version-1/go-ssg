package content

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/russross/blackfriday/v2"
)

func ConvertMarkdownToHTML(inputPath, outputDir string) {
	input, err := os.ReadFile(inputPath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", inputPath, err)
	}

	output := blackfriday.Run(input)

	outputFilePath := filepath.Join(outputDir, strings.TrimSuffix(filepath.Base(inputPath), ".md")+".html")
	err = os.WriteFile(outputFilePath, output, 0644)
	if err != nil {
		log.Fatalf("Failed to write file %s: %v", outputFilePath, err)
	}

	fmt.Printf("Converted %s to %s\n", inputPath, outputFilePath)
}
