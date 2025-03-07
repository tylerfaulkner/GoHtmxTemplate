package dotenv

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnv reads a .env file and sets environment variables.
func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("Error opening .env file:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines or comments.
		if len(line) == 0 || strings.HasPrefix(line, "#") {
			continue
		}

		// Split the line into key and value.
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue // or return an error if you want to be strict
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		log.Println("Setting env variable:", key, value)

		// Remove surrounding quotes if present.
		if len(value) > 1 && ((value[0] == '"' && value[len(value)-1] == '"') ||
			(value[0] == '\'' && value[len(value)-1] == '\'')) {
			value = value[1 : len(value)-1]
		}

		// Set the env variable.
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
