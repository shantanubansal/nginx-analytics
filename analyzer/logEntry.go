package analyzer

import (
	"regexp"
	"strconv"
)

type LogEntry struct {
	RemoteAddr                   string
	RemoteUser                   string
	TimeLocal                    string
	Request                      string
	Status                       int
	BodyBytesSent                int
	HttpReferer                  string
	HttpUserAgent                string
	RequestLength                int
	RequestTime                  float64
	ProxyUpstreamName            string
	ProxyAlternativeUpstreamName string
	UpstreamAddr                 string
	UpstreamResponseLength       int
	UpstreamResponseTime         float64
	UpstreamStatus               int
	ReqID                        string
	UpstreamHTTPSvcReqID         string
	UpstreamHTTPRequester        string
	UpstreamHTTPErrorCode        string
	UpstreamHTTPErrRef           string
}

func ParseLogEntry(logEntry string) LogEntry {
	pattern := `^(?P<remote_addr>[^ ]+) - (?P<remote_user>[^ ]+) \[(?P<time_local>[^\]]+)\] "(?P<request>[^"]+)" (?P<status>\d+) (?P<body_bytes_sent>\d+) "(?P<http_referer>[^"]*)" "(?P<http_user_agent>[^"]*)" (?P<request_length>\d+) (?P<request_time>[^ ]+) \[(?P<proxy_upstream_name>[^\]]*)\] \[(?P<proxy_alternative_upstream_name>[^\]]*)\] (?P<upstream_addr>[^ ]+) (?P<upstream_response_length>\d+) (?P<upstream_response_time>[^ ]+) (?P<upstream_status>\d+) (?P<req_id>[^ ]+) (?P<upstream_http_svc_req_id>[^ ]+) (?P<upstream_http_requester>[^ ]+) (?P<upstream_http_err_code>[^ ]+) (?P<upstream_http_err_ref>.+)$`

	re := regexp.MustCompile(pattern)
	match := re.FindStringSubmatch(logEntry)

	var le LogEntry
	for i, name := range re.SubexpNames() {
		if i > 0 && i <= len(match) {
			switch name {
			case "remote_addr":
				le.RemoteAddr = match[i]
			case "remote_user":
				le.RemoteUser = match[i]
			case "time_local":
				le.TimeLocal = match[i]
			case "request":
				le.Request = match[i]
			case "status":
				le.Status, _ = strconv.Atoi(match[i])
			case "body_bytes_sent":
				le.BodyBytesSent, _ = strconv.Atoi(match[i])
			case "http_referer":
				le.HttpReferer = match[i]
			case "http_user_agent":
				le.HttpUserAgent = match[i]
			case "request_length":
				le.RequestLength, _ = strconv.Atoi(match[i])
			case "request_time":
				le.RequestTime, _ = strconv.ParseFloat(match[i], 64)
			case "proxy_upstream_name":
				le.ProxyUpstreamName = match[i]
			case "proxy_alternative_upstream_name":
				le.ProxyAlternativeUpstreamName = match[i]
			case "upstream_addr":
				le.UpstreamAddr = match[i]
			case "upstream_response_length":
				le.UpstreamResponseLength, _ = strconv.Atoi(match[i])
			case "upstream_response_time":
				le.UpstreamResponseTime, _ = strconv.ParseFloat(match[i], 64)
			case "upstream_status":
				le.UpstreamStatus, _ = strconv.Atoi(match[i])
			case "req_id":
				le.ReqID = match[i]
			case "upstream_http_svc_req_id":
				le.UpstreamHTTPSvcReqID = match[i]
			case "upstream_http_requester":
				le.UpstreamHTTPRequester = match[i]
			case "upstream_http_err_code":
				le.UpstreamHTTPErrorCode = match[i]
			case "upstream_http_err_ref":
				le.UpstreamHTTPErrRef = match[i]
			}
		}
	}
	return le
}
