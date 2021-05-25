package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListKmsByTagsRequest struct {
	// 资源实例

	ResourceInstances string `json:"resource_instances"`
	// API版本号

	VersionId string `json:"version_id"`

	Body *ListKmsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListKmsByTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListKmsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListKmsByTagsRequest", string(data)}, " ")
}
