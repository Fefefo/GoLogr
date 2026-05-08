package main

import (
	"flag"
	"fmt"
	"log"

	"log-analyzer/internal/analyzer"
	"log-analyzer/internal/config"
	"log-analyzer/internal/formatter"
)

func main() {
	filePath := flag.String("file", "", "Path to log file")
	levelFilter := flag.String("level", "", "Filter by log level (e.g. ERROR)")
	topN := flag.Int("top", 5, "Number of top errors to show")
	jsonOutput := flag.Bool("json", false, "Output in JSON format")

	flag.Parse()

	if *filePath == "" {
		fmt.Println("Usage: log-analyzer --file <file.log>")
		return
	}

	cfg := config.Config{
		FilePath:    *filePath,
		LevelFilter: *levelFilter,
		TopN:        *topN,
		JsonOutput:  *jsonOutput,
	}

	result, err := analyzer.AnalyzeFile(cfg)
	if err != nil {
		log.Fatal(err)
	}

	sortedErrors := analyzer.GetSortedErrors(result.ErrorCount)
	limit := min(len(sortedErrors), cfg.TopN)

	if cfg.JsonOutput {
		formatter.PrintJSON(result, sortedErrors, limit)
	} else {
		formatter.PrintText(result, sortedErrors, limit)
	}
}
