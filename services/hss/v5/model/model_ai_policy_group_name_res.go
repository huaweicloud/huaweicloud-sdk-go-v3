package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupNameRes **参数解释**: 策略组名称 **取值范围**: 字符长度1-128位
type AiPolicyGroupNameRes struct {
}

func (o AiPolicyGroupNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupNameRes struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupNameRes", string(data)}, " ")
}
