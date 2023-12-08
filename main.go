package main

import (
	"encoding/csv"
	"fmt"
	"github.com/shantanubansal/nginx-analytics/analyzer"
	"log"
	"os"
	"strings"
)

var logEntries = []analyzer.LogEntry{}

func main() {
	err := processMessageColumn("input.csv")
	if err != nil {
		log.Fatal(err.Error())
	}
	analysisResults := analyzer.AnalyzeLogEntries(logEntries)
	err = analyzer.ExportToExcel(*analysisResults, "Result.xlsx")
	if err != nil {
		log.Fatal(err.Error())
	}

}

func processMessage(message string) {
	logEntries = append(logEntries, analyzer.ParseLogEntry(message))
}
func processMessageColumn(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	headers, err := csvReader.Read()
	if err != nil {
		return err
	}

	messageIndex := -1
	for i, header := range headers {
		if strings.ToLower(header) == "message" {
			messageIndex = i
			break
		}
	}

	if messageIndex == -1 {
		return fmt.Errorf("no 'Message' column found")
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return err
		}

		if len(record) > messageIndex {
			processMessage(record[messageIndex])
		}
	}

	return nil
}
