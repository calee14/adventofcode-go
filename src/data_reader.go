package src

import (
	"bufio"
	"os"
)

func ReadFile(file_path string) ([]string, error) {
	file, err := os.Open(file_path)
	if err != nil {
		return nil, err
	}
	defer func() {
		closeErr := file.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
