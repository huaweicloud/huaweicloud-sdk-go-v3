package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTagsRequest struct {

	// 资源类型。
	ResourceType *ResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`

	Body *CreateTagsRequestBody `json:"body,omitempty"`
}

func (o CreateTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateTagsRequest", string(data)}, " ")
}
