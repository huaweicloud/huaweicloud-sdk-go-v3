package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetRedisPitrPolicyRequest Request Object
type SetRedisPitrPolicyRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	Body *SetRedisPitrPolicyRequestBody `json:"body,omitempty"`
}

func (o SetRedisPitrPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetRedisPitrPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetRedisPitrPolicyRequest", string(data)}, " ")
}
