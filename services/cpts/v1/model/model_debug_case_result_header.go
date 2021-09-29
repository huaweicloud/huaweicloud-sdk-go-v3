package model

import (
	"encoding/json"

	"strings"
)

type DebugCaseResultHeader struct {
	// Connection

	Connection *string `json:"Connection,omitempty"`
	// Content-Type

	ContentType *string `json:"Content-Type,omitempty"`
	// Host

	Host *string `json:"Host,omitempty"`
}

func (o DebugCaseResultHeader) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DebugCaseResultHeader struct{}"
	}

	return strings.Join([]string{"DebugCaseResultHeader", string(data)}, " ")
}
