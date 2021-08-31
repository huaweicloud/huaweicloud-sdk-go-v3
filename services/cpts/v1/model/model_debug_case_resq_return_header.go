package model

import (
	"encoding/json"

	"strings"
)

// returnHeader
type DebugCaseResqReturnHeader struct {
	// Connection

	Connection *string `json:"Connection,omitempty"`
	// Content-Length

	ContentLength *string `json:"Content-Length,omitempty"`
	// Content-Type

	ContentType *string `json:"Content-Type,omitempty"`
	// Date

	Date *string `json:"Date,omitempty"`
	// Vary

	Vary *string `json:"Vary,omitempty"`
}

func (o DebugCaseResqReturnHeader) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugCaseResqReturnHeader struct{}"
	}

	return strings.Join([]string{"DebugCaseResqReturnHeader", string(data)}, " ")
}
