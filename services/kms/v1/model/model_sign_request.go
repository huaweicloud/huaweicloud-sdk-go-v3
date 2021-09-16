package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type SignRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *SignRequestBody `json:"body,omitempty"`
}

func (o SignRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SignRequest struct{}"
	}

	return strings.Join([]string{"SignRequest", string(data)}, " ")
}
