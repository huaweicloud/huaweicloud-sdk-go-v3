package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowResourceTagsRequest Request Object
type ShowResourceTagsRequest struct {

	// 资源类型，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ResourceType string `json:"resource_type"`

	// 资源id，获取方式请参见[如何获取接口URI中参数](cce_02_0271.xml)。
	ResourceId string `json:"resource_id"`
}

func (o ShowResourceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResourceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowResourceTagsRequest", string(data)}, " ")
}
