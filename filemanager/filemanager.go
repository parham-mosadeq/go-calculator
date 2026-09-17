package filemanager

import (
	"bufio"
	"encoding/json"
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

func WriteJSON(fileName string, data interface{}) error {
	file, err := os.Create(fileName)
	if err != nil {
		return errors.New("failed to create json")
	}

	err = json.NewEncoder(file).Encode(data)

	if err != nil {
		return errors.New("failed to convert to json")
	}

	file.Close()

	return nil
}
