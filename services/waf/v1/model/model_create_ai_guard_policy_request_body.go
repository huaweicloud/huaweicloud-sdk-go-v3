package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAiGuardPolicyRequestBody 创建策略请求
type CreateAiGuardPolicyRequestBody struct {

	// 策略名称
	Name string `json:"name"`
}

func (o CreateAiGuardPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAiGuardPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"CreateAiGuardPolicyRequestBody", string(data)}, " ")
}
