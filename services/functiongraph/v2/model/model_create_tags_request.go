package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateTagsRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 资源ID
	ResourceId string `json:"resource_id"`

	Body *UpdateFunctionTagsRequestBody `json:"body,omitempty"`
}

func (o CreateTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateTagsRequest", string(data)}, " ")
}
