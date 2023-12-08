package analyzer

import (
	"fmt"
	"github.com/xuri/excelize/v2"
)

func ExportToExcel(results AnalysisResults, fileName string) error {
	f := excelize.NewFile()
	createRepetitiveRequestCountSheet(f, results)
	createRepetitiveRequestStatusCountSheet(f, results)
	createRepetitiveRequestErrorCodeCountSheet(f, results)
	createRepetitiveRequestErrorCodeUserCountSheet(f, results)
	createRequestTimeStatsSheet(f, results)
	if err := f.SaveAs(fileName); err != nil {
		return fmt.Errorf("failed to save excel file: %v", err)
	}
	return nil
}

func createRepetitiveRequestCountSheet(f *excelize.File, results AnalysisResults) {
	sheetName := "Repetitive Request Count"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Total Repetitive Requests")
	f.SetCellValue(sheetName, "B1", results.RepetitiveRequestCount)
}

func createRepetitiveRequestStatusCountSheet(f *excelize.File, results AnalysisResults) {
	sheetName := "Repetitive Request Status Count"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Request + Status")
	f.SetCellValue(sheetName, "B1", "Count")

	row := 2
	for key, count := range results.RepetitiveRequestStatusCount {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), key)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), count)
		row++
	}
}

func createRepetitiveRequestErrorCodeCountSheet(f *excelize.File, results AnalysisResults) {
	sheetName := "Repetitive Request Error Code Count"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Request + Error Code")
	f.SetCellValue(sheetName, "B1", "Count")

	row := 2
	for key, count := range results.RepetitiveRequestErrorCodeCount {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), key)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), count)
		row++
	}
}

func createRepetitiveRequestErrorCodeUserCountSheet(f *excelize.File, results AnalysisResults) {
	sheetName := "Repetitive Request Error Code User Count"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Request + Error Code + User")
	f.SetCellValue(sheetName, "B1", "Count")

	row := 2
	for key, count := range results.RepetitiveRequestErrorCodeUserCount {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), key)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), count)
		row++
	}
}

func createRequestTimeStatsSheet(f *excelize.File, results AnalysisResults) {
	sheetName := "Request Time Stats"
	f.NewSheet(sheetName)
	f.SetCellValue(sheetName, "A1", "Request")
	f.SetCellValue(sheetName, "B1", "Max Time")
	f.SetCellValue(sheetName, "C1", "Mean Time")
	f.SetCellValue(sheetName, "D1", "Min Time")

	row := 2
	for request, stats := range results.RequestTimeStats {
		f.SetCellValue(sheetName, fmt.Sprintf("A%d", row), request)
		f.SetCellValue(sheetName, fmt.Sprintf("B%d", row), stats.MaxTime)
		f.SetCellValue(sheetName, fmt.Sprintf("C%d", row), stats.MeanTime)
		f.SetCellValue(sheetName, fmt.Sprintf("D%d", row), stats.MinTime)
		row++
	}
}
