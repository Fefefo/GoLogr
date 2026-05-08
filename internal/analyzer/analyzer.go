package analyzer

import (
	"bufio"
	"os"
	"sort"

	"log-analyzer/internal/config"
	"log-analyzer/internal/parser"
)

type Result struct {
	TotalLines int

	LevelCount map[string]int
	ErrorCount map[string]int
}

type KV struct {
	Key   string
	Value int
}

func AnalyzeFile(cfg config.Config) (*Result, error) {
	file, err := os.Open(cfg.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	levelCount := make(map[string]int)
	errorCount := make(map[string]int)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()
		count++

		level := parser.ParseLevel(line)
		if cfg.LevelFilter != "" && level != cfg.LevelFilter {
			continue
		}

		levelCount[level]++

		if level == "ERROR" {
			msg := parser.ParseErrorMessage(line)
			errorCount[msg]++
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &Result{
		TotalLines: count,
		LevelCount: levelCount,
		ErrorCount: errorCount,
	}, nil
}

func GetSortedErrors(errorMap map[string]int) []KV {
	var sortedErrors []KV

	for k, v := range errorMap {
		sortedErrors = append(sortedErrors, KV{k, v})
	}

	sort.Slice(sortedErrors, func(i, j int) bool {
		return sortedErrors[i].Value > sortedErrors[j].Value
	})
	return sortedErrors
}
