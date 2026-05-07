package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleTypeReqM **参数解释**： 规则类型 **约束限制**： 必填 **取值范围**： - black_hash：黑hash  **默认取值**： 不涉及
type RuleTypeReqM struct {
}

func (o RuleTypeReqM) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleTypeReqM struct{}"
	}

	return strings.Join([]string{"RuleTypeReqM", string(data)}, " ")
}
