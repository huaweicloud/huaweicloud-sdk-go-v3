package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupIdRes **参数解释**： 策略组ID **取值范围**： 字符长度1-20位
type AiPolicyGroupIdRes struct {
}

func (o AiPolicyGroupIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupIdRes struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupIdRes", string(data)}, " ")
}
