package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type InputTarget struct {
	Id   string `json:"id"`
	Path string `json:"path"`
	Type string `json:"type"`
}

type OutputTarget struct {
	FilePath     string `json:"file_path"`
	Status       string `json:"status"`
	Message      string `json:"message"`
	ErrorDetails string `json:"error_details"`
}

func LoadTargetsFromFile(filepath string) ([]InputTarget, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read files %s : %w", filepath, err)
	}
	var targets []InputTarget
	if err := json.Unmarshal(data, &targets); err != nil {
		return nil, fmt.Errorf("failed to unmarshal %s : %w ", filepath, err)
	}
	return targets, nil
}
