package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteAiGuardPolicyRequestBody 批量删除策略请求
type BatchDeleteAiGuardPolicyRequestBody struct {

	// 策略 ID 列表
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchDeleteAiGuardPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteAiGuardPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"BatchDeleteAiGuardPolicyRequestBody", string(data)}, " ")
}
