package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowKmsTagsRequest struct {
	// API版本号

	VersionId string `json:"version_id"`
	// 密钥ID

	KeyId string `json:"key_id"`
}

func (o ShowKmsTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowKmsTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowKmsTagsRequest", string(data)}, " ")
}
