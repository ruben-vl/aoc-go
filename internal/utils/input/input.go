package input

import (
	"bufio"
	"fmt"
	"io"
)

func ReadSingleLine(r io.Reader) (string, error) {
	scanner := bufio.NewScanner(r)
	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return "", fmt.Errorf("reading input: %w", err)
		}
		return "", fmt.Errorf("no input line found")
	}
	return scanner.Text(), nil
}
