package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateInstanceConsumerGroupRequest Request Object
type UpdateInstanceConsumerGroupRequest struct {

	// 引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组ID。
	Group string `json:"group"`

	Body *GroupCreateReq `json:"body,omitempty"`
}

func (o UpdateInstanceConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceConsumerGroupRequest", string(data)}, " ")
}
