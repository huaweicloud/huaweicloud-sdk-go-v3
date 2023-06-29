package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagsRequest Request Object
type ShowResourceTagsRequest struct {

	// 资源类型。
	ResourceType *ResourceType `json:"resource_type"`

	// 资源ID。
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsRequest", string(data)}, " ")
}
