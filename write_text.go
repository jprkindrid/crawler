package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
)

func writeBodyToText(normalizedURL, htmlBody string) error {
	err := os.MkdirAll("pages", 0755)
	if err != nil {
		return fmt.Errorf("failed to create pages directory: %w", err)
	}

	writeBody := normalizedURL + "\n\n" + htmlBody

	hash := sha256.Sum256([]byte(normalizedURL))
	filename := hex.EncodeToString(hash[:]) + ".html"
	filePath := filepath.Join("pages", filename)

	err = os.WriteFile(filePath, []byte(writeBody), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file %s: %w", filePath, err)
	}

	return nil
}
