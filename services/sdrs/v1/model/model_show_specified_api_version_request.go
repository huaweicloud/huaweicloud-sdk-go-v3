package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSpecifiedApiVersionRequest struct {
	// API版本号。例如: v1。

	ApiVersion string `json:"api_version"`
}

func (o ShowSpecifiedApiVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSpecifiedApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowSpecifiedApiVersionRequest", string(data)}, " ")
}
