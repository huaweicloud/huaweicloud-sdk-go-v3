package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteGcbResourceTagsRequest Request Object
type BatchDeleteGcbResourceTagsRequest struct {

	// 资源唯一标识符。
	ResourceId string `json:"resource_id"`

	Body *CreateDeleteGcbTagsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteGcbResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteGcbResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteGcbResourceTagsRequest", string(data)}, " ")
}
