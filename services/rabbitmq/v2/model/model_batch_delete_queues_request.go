package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteQueuesRequest Request Object
type BatchDeleteQueuesRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	Body *BatchDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteQueuesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteQueuesRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteQueuesRequest", string(data)}, " ")
}
