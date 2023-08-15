package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartitionEndMessageRequest Request Object
type ShowPartitionEndMessageRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// Topic名称。  Topic名称必现以字母开头且只支持大小写字母、中横线、下划线以及数字。
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
