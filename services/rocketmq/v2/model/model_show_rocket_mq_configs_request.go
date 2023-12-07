package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRocketMqConfigsRequest Request Object
type ShowRocketMqConfigsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`
}

func (o ShowRocketMqConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRocketMqConfigsRequest struct{}"
	}

	return strings.Join([]string{"ShowRocketMqConfigsRequest", string(data)}, " ")
}
