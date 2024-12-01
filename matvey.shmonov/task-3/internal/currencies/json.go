package currencies

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func MarshalJSON(c *Currencies, f string) error {
	jsonData, err := json.MarshalIndent(c.List, "", "  ")
	if err != nil {
		return fmt.Errorf("unable to convert to json: %v", err)
	}
	dir := filepath.Dir(f)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("unable to create directories: %v", err)
	}

	file, err := os.Create(f)
	if err != nil {
		return fmt.Errorf("unable to create file: %v", err)
	}
	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		return fmt.Errorf("unable to write to file: %v", err)
	}

	return nil
}
