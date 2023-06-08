package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type SendDlqMessageRequest struct {

	// engine
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *DeadletterResendReq `json:"body,omitempty"`
}

func (o SendDlqMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageRequest struct{}"
	}

	return strings.Join([]string{"SendDlqMessageRequest", string(data)}, " ")
}
