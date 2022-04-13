package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePartitionRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
	// Topic名称。

	Topic string `json:"topic"`

	Body *CreatePartitionReq `json:"body,omitempty"`
}

func (o CreatePartitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionRequest struct{}"
	}

	return strings.Join([]string{"CreatePartitionRequest", string(data)}, " ")
}
