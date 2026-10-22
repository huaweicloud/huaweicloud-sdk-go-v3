package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAiGuardPolicyRequestBody 更新策略请求
type UpdateAiGuardPolicyRequestBody struct {

	// 策略名称
	Name string `json:"name"`
}

func (o UpdateAiGuardPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAiGuardPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAiGuardPolicyRequestBody", string(data)}, " ")
}
