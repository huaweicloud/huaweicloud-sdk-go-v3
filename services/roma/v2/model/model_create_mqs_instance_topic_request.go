package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateMqsInstanceTopicRequest struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`

	Body *CreateMqsInstanceTopicReq `json:"body,omitempty"`
}

func (o CreateMqsInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMqsInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"CreateMqsInstanceTopicRequest", string(data)}, " ")
}
