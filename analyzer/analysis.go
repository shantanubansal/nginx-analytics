package analyzer

import (
	"strconv"
)

type AnalysisResults struct {
	RepetitiveRequestCount              int
	RepetitiveRequestStatusCount        map[string]int
	RepetitiveRequestErrorCodeCount     map[string]int
	RepetitiveRequestErrorCodeUserCount map[string]int
	RequestTimeStats                    map[string]RequestTimeStat
}

type RequestTimeStat struct {
	MaxTime   float64
	MeanTime  float64
	MinTime   float64
	TotalTime float64
	Count     int
}

func AnalyzeLogEntries(logEntries []LogEntry) *AnalysisResults {
	results := &AnalysisResults{
		RepetitiveRequestStatusCount:        make(map[string]int),
		RepetitiveRequestErrorCodeCount:     make(map[string]int),
		RepetitiveRequestErrorCodeUserCount: make(map[string]int),
		RequestTimeStats:                    make(map[string]RequestTimeStat),
	}

	requestCount := make(map[string]int)
	for _, entry := range logEntries {
		requestKey := entry.Request
		requestStatusKey := entry.Request + strconv.Itoa(entry.Status)
		requestErrorCodeKey := entry.Request + entry.UpstreamHTTPErrorCode
		requestErrorCodeUserKey := entry.Request + entry.UpstreamHTTPErrorCode + entry.UpstreamHTTPRequester

		// Count repetitive requests
		requestCount[requestKey]++
		if requestCount[requestKey] > 1 {
			results.RepetitiveRequestCount++
		}

		// Count combinations of repetitive requests
		results.RepetitiveRequestStatusCount[requestStatusKey]++
		results.RepetitiveRequestErrorCodeCount[requestErrorCodeKey]++
		results.RepetitiveRequestErrorCodeUserCount[requestErrorCodeUserKey]++

		// Track request time statistics
		timeStat := results.RequestTimeStats[entry.Request]
		timeStat.TotalTime += entry.UpstreamResponseTime
		timeStat.Count++
		if entry.UpstreamResponseTime > timeStat.MaxTime {
			timeStat.MaxTime = entry.UpstreamResponseTime
		}
		if entry.UpstreamResponseTime < timeStat.MinTime || timeStat.MinTime == 0 {
			timeStat.MinTime = entry.UpstreamResponseTime
		}
		results.RequestTimeStats[entry.Request] = timeStat
	}

	// Calculate mean time for each request
	for request, stat := range results.RequestTimeStats {
		if stat.Count > 0 {
			stat.MeanTime = stat.TotalTime / float64(stat.Count)
		} else {
			stat.MeanTime = 0
		}
		results.RequestTimeStats[request] = stat
	}

	return results
}
