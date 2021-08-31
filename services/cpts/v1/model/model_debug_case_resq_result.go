package model

import (
	"encoding/json"

	"strings"
)

type DebugCaseResqResult struct {
	// body

	Body *string `json:"body,omitempty"`
	// errorReason

	ErrorReason *string `json:"errorReason,omitempty"`

	Header *DebugCaseResqHeader `json:"header,omitempty"`
	// name

	Name *string `json:"name,omitempty"`
	// responseTime

	ResponseTime *int32 `json:"responseTime,omitempty"`
	// result

	Result *int32 `json:"result,omitempty"`
	// returnBody

	ReturnBody *string `json:"returnBody,omitempty"`

	ReturnHeader *DebugCaseResqReturnHeader `json:"returnHeader,omitempty"`
	// statusCode

	StatusCode *string `json:"statusCode,omitempty"`
	// url

	Url *string `json:"url,omitempty"`
}

func (o DebugCaseResqResult) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugCaseResqResult struct{}"
	}

	return strings.Join([]string{"DebugCaseResqResult", string(data)}, " ")
}
