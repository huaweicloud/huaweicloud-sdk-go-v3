package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyIdRes **参数解释**: 策略ID **取值范围**: 字符长度1-20位
type AiPolicyIdRes struct {
}

func (o AiPolicyIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyIdRes struct{}"
	}

	return strings.Join([]string{"AiPolicyIdRes", string(data)}, " ")
}
