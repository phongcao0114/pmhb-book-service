package utils

import (
	"time"
)

type (
	// RequestInfo struct contains request info
	RequestInfo struct {
		Body interface{} `json:"request_body"`
	}

	// KbankResponseHeader struct contains response header
	KbankResponseHeader struct {
		ResponseAppID string                   `json:"response_app_id"`
		ResponseDate  time.Time                `json:"response_datetime"`
		StatusCode    string                   `json:"status_code"`
		Errors        ResponseErrorKbankHeader `json:"error"`
	}

	// ResponseErrorKbankHeader struct contains response error header
	ResponseErrorKbankHeader struct {
		ErrorCode        string `json:"error_code"`
		ErrorDesc        string `json:"error_desc"`
		OriginalErrorStr string `json:"original_error_str"`
	}
)
