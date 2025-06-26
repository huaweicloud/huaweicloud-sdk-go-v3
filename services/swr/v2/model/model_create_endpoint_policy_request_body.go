package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateEndpointPolicyRequestBody struct {

	// true为开启，false为关闭。只有Disable、EnableFailed开启失败或者关闭时才能开启，只有Enable、DisableFailed时才能关闭。
	Enable bool `json:"enable"`
}

func (o CreateEndpointPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEndpointPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateEndpointPolicyRequestBody", string(data)}, " ")
}
