package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelSelfGrantRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *RevokeGrantRequestBody `json:"body,omitempty"`
}

func (o CancelSelfGrantRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelSelfGrantRequest struct{}"
	}

	return strings.Join([]string{"CancelSelfGrantRequest", string(data)}, " ")
}
