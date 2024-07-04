package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueueRequest Request Object
type CreateQueueRequest struct {

	// 实例ID
	InstanceId string `json:"instance_id"`

	// Vhost名称
	Vhost string `json:"vhost"`

	Body *CreateQueueBody `json:"body,omitempty"`
}

func (o CreateQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueRequest struct{}"
	}

	return strings.Join([]string{"CreateQueueRequest", string(data)}, " ")
}
