package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CancelKeyDeletionRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o CancelKeyDeletionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CancelKeyDeletionRequest struct{}"
	}

	return strings.Join([]string{"CancelKeyDeletionRequest", string(data)}, " ")
}
