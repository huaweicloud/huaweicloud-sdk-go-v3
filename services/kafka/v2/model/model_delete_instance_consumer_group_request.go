package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInstanceConsumerGroupRequest Request Object
type DeleteInstanceConsumerGroupRequest struct {

	// 引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组ID。
	Group string `json:"group"`
}

func (o DeleteInstanceConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInstanceConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteInstanceConsumerGroupRequest", string(data)}, " ")
}
