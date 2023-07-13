package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketmqTagsRequest Request Object
type ShowRocketmqTagsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowRocketmqTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketmqTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowRocketmqTagsRequest", string(data)}, " ")
}
