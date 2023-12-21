package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceTagsRequest Request Object
type ShowInstanceTagsRequest struct {

	// 资源ID。(list接口获取)
	ResourceId string `json:"resource_id"`
}

func (o ShowInstanceTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowInstanceTagsRequest", string(data)}, " ")
}
