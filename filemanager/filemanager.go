package filemanager

import (
	"bufio"
	"errors"
	"os"
)

func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, errors.New("Reading file failed!")
	}

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {

		file.Close()
		return nil, errors.New("Reading the file content failed!")
	}

	return lines, nil
}
