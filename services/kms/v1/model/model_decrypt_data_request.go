package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DecryptDataRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *DecryptDataRequestBody `json:"body,omitempty"`
}

func (o DecryptDataRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DecryptDataRequest struct{}"
	}

	return strings.Join([]string{"DecryptDataRequest", string(data)}, " ")
}
