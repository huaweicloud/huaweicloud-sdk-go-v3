package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListKeyDetailRequest struct {
	// API版本号

	VersionId string `json:"version_id"`

	Body *OperateKeyRequestBody `json:"body,omitempty"`
}

func (o ListKeyDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListKeyDetailRequest struct{}"
	}

	return strings.Join([]string{"ListKeyDetailRequest", string(data)}, " ")
}
