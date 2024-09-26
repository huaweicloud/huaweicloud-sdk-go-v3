package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateGcbResourceTagsRequest Request Object
type BatchCreateGcbResourceTagsRequest struct {

	// 资源唯一标识符。
	ResourceId string `json:"resource_id"`

	Body *CreateDeleteGcbTagsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateGcbResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateGcbResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateGcbResourceTagsRequest", string(data)}, " ")
}
