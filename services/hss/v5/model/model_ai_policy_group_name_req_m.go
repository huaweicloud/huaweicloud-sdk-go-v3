package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupNameReqM **参数解释**： 策略组名称 **约束限制**： 必填 **取值范围**： 字符长度1-128位 **默认取值**： 不涉及
type AiPolicyGroupNameReqM struct {
}

func (o AiPolicyGroupNameReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupNameReqM struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupNameReqM", string(data)}, " ")
}
