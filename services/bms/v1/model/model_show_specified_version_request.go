package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowSpecifiedVersionRequest struct {
	// API版本号

	ApiVersion string `json:"api_version"`
}

func (o ShowSpecifiedVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowSpecifiedVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowSpecifiedVersionRequest", string(data)}, " ")
}
