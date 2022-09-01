package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchUpdateConsumerGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id" xml:"instance_id"`

	Body *BatchUpdateConsumerGroupReq `json:"body,omitempty" xml:"body"`
}

func (o BatchUpdateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupRequest", string(data)}, " ")
}
