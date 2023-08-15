package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateTagsRequest Request Object
type CreateTagsRequest struct {

	// 资源类型
	ResourceType string `json:"resource_type"`

	// 资源id
	ResourceId string `json:"resource_id"`

	Body *CreateTagsReqbody `json:"body,omitempty"`
}

func (o CreateTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTagsRequest struct{}"
	}

	return strings.Join([]string{"CreateTagsRequest", string(data)}, " ")
}
