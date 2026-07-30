package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyDescriptionRes **参数解释**: 策略描述 **取值范围**: 字符长度1-256位
type AiPolicyDescriptionRes struct {
}

func (o AiPolicyDescriptionRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyDescriptionRes struct{}"
	}

	return strings.Join([]string{"AiPolicyDescriptionRes", string(data)}, " ")
}
