package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteConsumerGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	Group string `json:"group"`
}

func (o DeleteConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteConsumerGroupRequest", string(data)}, " ")
}
