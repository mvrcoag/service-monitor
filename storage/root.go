package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

var filePath = "./.service-monitor.json"

type Storage struct {
	Urls []string `json:"urls"`
}

func InitStorage() {
	_, err := os.Stat(filePath)

	if !os.IsNotExist(err) {
		return
	}

	_, err = os.Create(filePath)

	if err != nil {
		fmt.Println("Error creating storage")
		os.Exit(1)
	}

	s := Storage{
		Urls: []string{},
	}

	WriteStorage(&s)
}

func ReadStorage(storage *Storage) {
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error reading from storage")
		os.Exit(1)
	}

	defer file.Close()

	bytes, _ := io.ReadAll(file)

	err = json.Unmarshal(bytes, storage)

	if err != nil {
		fmt.Println("Error parsing JSON from storage")
		os.Exit(1)
	}
}

func WriteStorage(storage *Storage) {
	bytes, err := json.Marshal(storage)

	if err != nil {
		fmt.Println("Error encoding storage to JSON")
		os.Exit(1)
	}

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, os.ModeAppend)

	if err != nil {
		fmt.Println("Error reading from storage")
		os.Exit(1)
	}

	defer file.Close()

	_, err = file.Write(bytes)

	if err != nil {
		fmt.Println("Error writing to storage")
		os.Exit(1)
	}
}
