package analyzer

import (
	"os"
	"testing"

	"log-analyzer/internal/config"
)

func TestAnalyzeFile(t *testing.T) {
	content := `2026-04-23 INFO Start
2026-04-23 ERROR DB failed
2026-04-23 ERROR DB failed
2026-04-23 INFO End`

	tmpFile, err := os.CreateTemp("", "test.log")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.WriteString(content)
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	cfg := config.Config{
		FilePath: tmpFile.Name(),
	}

	result, err := AnalyzeFile(cfg)
	if err != nil {
		t.Fatal(err)
	}

	if result.TotalLines != 4 {
		t.Errorf("expected 4 lines, got %d", result.TotalLines)
	}

	if result.LevelCount["ERROR"] != 2 {
		t.Errorf("expected 2 ERROR, got %d", result.LevelCount["ERROR"])
	}

	if result.ErrorCount["DB failed"] != 2 {
		t.Errorf("expected 2 DB failed, got %d", result.ErrorCount["DB failed"])
	}
}
