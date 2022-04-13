package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListKmsByTagsRequest struct {
	// 资源实例，固定值为resource_instances

	ResourceInstances string `json:"resource_instances"`

	Body *ListKmsByTagsRequestBody `json:"body,omitempty"`
}

func (o ListKmsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListKmsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListKmsByTagsRequest", string(data)}, " ")
}
