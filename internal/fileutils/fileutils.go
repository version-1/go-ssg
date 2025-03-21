package fileutils

import (
	"fmt"
	"os"
)

func WriteHTMLToFile(filePath string, data []byte) error {
	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %v", filePath, err)
	}
	fmt.Printf("Converted to %s\n", filePath)
	return nil
}
