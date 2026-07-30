package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiPolicyGroupIdReqM **参数解释**： 策略组ID **约束限制**： 必填 **取值范围**： 字符长度1-20位 **默认取值**： 不涉及
type AiPolicyGroupIdReqM struct {
}

func (o AiPolicyGroupIdReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiPolicyGroupIdReqM struct{}"
	}

	return strings.Join([]string{"AiPolicyGroupIdReqM", string(data)}, " ")
}
