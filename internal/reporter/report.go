package reporter

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/docblizzard/loganizer/internal/config"
)

func ExportResultsToJsonfile(results []config.OutputTarget) error {
	data, err := json.MarshalIndent(results, "", "")
	if err != nil {
		return fmt.Errorf("Failed to marshal %s", err)
	}
	if err := os.WriteFile("report/report.json", data, 0600); err != nil {
		return fmt.Errorf("Failed to write %s", err)
	}
	return nil
}
