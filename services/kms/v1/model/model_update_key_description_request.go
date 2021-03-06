package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateKeyDescriptionRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *UpdateKeyDescriptionRequestBody `json:"body,omitempty"`
}

func (o UpdateKeyDescriptionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateKeyDescriptionRequest struct{}"
	}

	return strings.Join([]string{"UpdateKeyDescriptionRequest", string(data)}, " ")
}
