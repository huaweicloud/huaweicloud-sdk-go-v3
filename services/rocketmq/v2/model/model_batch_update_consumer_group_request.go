package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateConsumerGroupRequest Request Object
type BatchUpdateConsumerGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *BatchUpdateConsumerGroupReq `json:"body,omitempty"`
}

func (o BatchUpdateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroupRequest", string(data)}, " ")
}
