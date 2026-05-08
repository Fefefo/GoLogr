package formatter

import (
	"encoding/json"
	"fmt"
	"log-analyzer/internal/analyzer"
)

type JSONError struct {
	Message string `json:"message"`
	Count   int    `json:"count"`
}

type JSONOutput struct {
	TotalLines int            `json:"total_lines"`
	Levels     map[string]int `json:"levels"`
	TopErrors  []JSONError    `json:"top_errors"`
}

func PrintText(result *analyzer.Result, topErrors []analyzer.KV, limit int) {
	fmt.Println("Total lines:", result.TotalLines)

	fmt.Println("\nLevels:")
	for level, c := range result.LevelCount {
		fmt.Printf("%s: %d\n", level, c)
	}

	fmt.Println("\nTop Errors:")
	for i := range limit {
		fmt.Printf("%s: %d\n", topErrors[i].Key, topErrors[i].Value)
	}
}

func PrintJSON(result *analyzer.Result, topErrors []analyzer.KV, limit int) {
	var errors []JSONError

	for i := range limit {
		errors = append(errors, JSONError{
			Message: topErrors[i].Key,
			Count:   topErrors[i].Value,
		})
	}

	output := JSONOutput{
		TotalLines: result.TotalLines,
		Levels:     result.LevelCount,
		TopErrors:  errors,
	}

	data, err := json.MarshalIndent(output, "", "  ")
	if err != nil {
		fmt.Println("Error generating JSON:", err)
		return
	}

	fmt.Println(string(data))
}
