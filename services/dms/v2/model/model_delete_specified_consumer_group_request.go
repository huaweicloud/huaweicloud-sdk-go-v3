package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSpecifiedConsumerGroupRequest struct {
	// 队列ID

	QueueId string `json:"queue_id"`
	// 待删除的消费组ID

	GroupId string `json:"group_id"`
}

func (o DeleteSpecifiedConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSpecifiedConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteSpecifiedConsumerGroupRequest", string(data)}, " ")
}
