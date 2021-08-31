package model

import (
	"encoding/json"

	"strings"
)

// header
type DebugCaseResqHeader struct {
	// Connection

	Connection *string `json:"Connection,omitempty"`
	// Content-Type

	ContentType *string `json:"Content-Type,omitempty"`
	// Host

	Host *string `json:"Host,omitempty"`
}

func (o DebugCaseResqHeader) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugCaseResqHeader struct{}"
	}

	return strings.Join([]string{"DebugCaseResqHeader", string(data)}, " ")
}
