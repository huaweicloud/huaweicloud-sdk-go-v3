package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowPartitionEndMessageRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// Topic名称。

	Topic string `json:"topic"`
	// 分区编号。

	Partition int32 `json:"partition"`
}

func (o ShowPartitionEndMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionEndMessageRequest struct{}"
	}

	return strings.Join([]string{"ShowPartitionEndMessageRequest", string(data)}, " ")
}
