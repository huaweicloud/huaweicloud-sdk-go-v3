package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DecryptDatakeyRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *DecryptDatakeyRequestBody `json:"body,omitempty"`
}

func (o DecryptDatakeyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DecryptDatakeyRequest struct{}"
	}

	return strings.Join([]string{"DecryptDatakeyRequest", string(data)}, " ")
}
