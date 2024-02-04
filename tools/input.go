package tools

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(fileName string) (<-chan string, error) {
	line := make(chan string)

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}

	scanner := bufio.NewScanner(file)

	go func() {
		defer close(line)
		defer file.Close()

		for scanner.Scan() {
			line <- scanner.Text()
		}
	}()

	return line, nil
}
