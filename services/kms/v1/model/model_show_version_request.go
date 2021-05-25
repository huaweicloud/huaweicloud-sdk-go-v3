package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowVersionRequest struct {
	// API版本号

	VersionId string `json:"version_id"`
}

func (o ShowVersionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVersionRequest struct{}"
	}

	return strings.Join([]string{"ShowVersionRequest", string(data)}, " ")
}
