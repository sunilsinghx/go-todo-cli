package main

import (
	"encoding/json"
	"os"
)

// Storage is a generic type that represents a storage system for saving and loading data.
type Storage[T any] struct {
	FileName string // The name of the file where data will be saved or loaded from.
}

// NewStorage creates a new Storage instance with the specified file name.
func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{FileName: fileName}
}

// Save serializes the provided data and writes it to the specified file.
func (s *Storage[T]) Save(data T) error {
	// Marshal the data into a JSON format with indentation for readability.
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		// Return error if marshaling fails.
		return err
	}

	// Write the marshaled data to the file.
	return os.WriteFile(s.FileName, fileData, 0644)
}

// Load reads data from the specified file and deserializes it into the provided variable.
func (s *Storage[T]) Load(data *T) error {
	// Read the raw data from the file.
	fileData, err := os.ReadFile(s.FileName)
	if err != nil {
		// Return error if reading the file fails.
		return err
	}

	// Unmarshal the JSON data into the provided variable.
	return json.Unmarshal(fileData, data)
}
