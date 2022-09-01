package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreatePartitionRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	// Topic名称。
	Topic string `json:"topic" xml:"topic"`

	Body *CreatePartitionReq `json:"body,omitempty" xml:"body"`
}

func (o CreatePartitionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePartitionRequest struct{}"
	}

	return strings.Join([]string{"CreatePartitionRequest", string(data)}, " ")
}
