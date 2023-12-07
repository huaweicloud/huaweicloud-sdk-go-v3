package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRocketMqConfigsRequest Request Object
type UpdateRocketMqConfigsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ModifyConfigReq `json:"body,omitempty"`
}

func (o UpdateRocketMqConfigsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRocketMqConfigsRequest struct{}"
	}

	return strings.Join([]string{"UpdateRocketMqConfigsRequest", string(data)}, " ")
}
