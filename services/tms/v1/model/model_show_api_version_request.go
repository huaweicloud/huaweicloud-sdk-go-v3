package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowApiVersionRequest struct {
	// API版本号。

	ApiVersion string `json:"api_version"`
}

func (o ShowApiVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApiVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowApiVersionRequest", string(data)}, " ")
}
