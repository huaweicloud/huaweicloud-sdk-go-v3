package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SendRocketMqDlqMessageRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeadletterResendReq `json:"body,omitempty"`
}

func (o SendRocketMqDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendRocketMqDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"SendRocketMqDlqMessageRequest", string(data)}, " ")
}
