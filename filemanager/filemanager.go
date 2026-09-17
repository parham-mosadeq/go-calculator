package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManager struct {
	InputFilepath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilepath)

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

func (fm FileManager) WriteJSON(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
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

func New(InputFilepath, OutputFilePath string) FileManager {
	return FileManager{InputFilepath: InputFilepath, OutputFilePath: OutputFilePath}
}
