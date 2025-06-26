package checker

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/docblizzard/loganizer/internal/config"
)

var keywords = []string{"error", "invalid", "fail", "exception"}

func ParseLog(target config.InputTarget) config.OutputTarget {
	var output config.OutputTarget
	output.FilePath = target.Path

	file, err := os.Open(target.Path)
	if err != nil {
		return config.OutputTarget{
			Id:           "invalid-path",
			FilePath:     target.Path,
			Status:       "FAILED",
			Message:      "Fichier introuvable.",
			ErrorDetails: err.Error(),
		}
	}
	defer file.Close()

	var matchedLines []string
	scanner := bufio.NewScanner(file)
	lineNumber := 1

	for scanner.Scan() {
		line := scanner.Text()
		for _, keyword := range keywords {
			if strings.Contains(strings.ToLower(line), keyword) {
				matchedLines = append(matchedLines, fmt.Sprintf("Ligne %d: %s", lineNumber, line))
				break
			}
		}
		lineNumber++
	}

	if len(matchedLines) == 0 {
		output.Id = target.Id
		output.Status = "ok"
		output.Message = "Aucune erreur détectée"
		output.ErrorDetails = ""
	} else {
		output.Status = "warning"
		output.Message = fmt.Sprintf("%d ligne(s) suspecte(s) détectée(s)", len(matchedLines))
		output.ErrorDetails = strings.Join(matchedLines, "\n")
	}

	return output
}
