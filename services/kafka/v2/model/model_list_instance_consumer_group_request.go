package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupRequest Request Object
type ListInstanceConsumerGroupRequest struct {

	// 引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名过滤查询。
	Group string `json:"group"`
}

func (o ListInstanceConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupRequest", string(data)}, " ")
}
