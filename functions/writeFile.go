package functions

import (
	"io"
	"os"
	"strings"
)

func WriteFile(filename string, secretsJson map[string]interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for key, value := range secretsJson {
		s := []string{key, value.(string)}
		_, err = io.WriteString(file, strings.Join(s, "=") + "\n")
	}

	if err != nil {
		return err
	}

	return file.Sync()
}
