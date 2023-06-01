package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ValidateRocketMqConsumedMessageRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *ResendReq `json:"body,omitempty"`
}

func (o ValidateRocketMqConsumedMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ValidateRocketMqConsumedMessageRequest struct{}"
	}

	return strings.Join([]string{"ValidateRocketMqConsumedMessageRequest", string(data)}, " ")
}
