package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateMqsInstanceTopicRequest Request Object
type UpdateMqsInstanceTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *UpdateMqsInstanceTopicReq `json:"body,omitempty"`
}

func (o UpdateMqsInstanceTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateMqsInstanceTopicRequest struct{}"
	}

	return strings.Join([]string{"UpdateMqsInstanceTopicRequest", string(data)}, " ")
}
