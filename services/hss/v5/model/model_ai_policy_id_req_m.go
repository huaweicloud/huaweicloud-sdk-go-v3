package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyIdReqM **参数解释**: 策略ID **约束限制**: 必填 **取值范围**: 字符长度1-20位 **默认取值**: 不涉及
type AiPolicyIdReqM struct {
}

func (o AiPolicyIdReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyIdReqM struct{}"
	}

	return strings.Join([]string{"AiPolicyIdReqM", string(data)}, " ")
}
