package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateKeyRotationIntervalRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *UpdateKeyRotationIntervalRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyRotationIntervalRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateKeyRotationIntervalRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyRotationIntervalRequest", string(data)}, " ")
}
