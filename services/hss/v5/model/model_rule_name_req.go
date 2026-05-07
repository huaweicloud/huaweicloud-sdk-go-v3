package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleNameReq **参数解释**： 规则名称 **约束限制**： 不涉及 **取值范围**： 字符长度1-64位 **默认取值**： 不涉及
type RuleNameReq struct {
}

func (o RuleNameReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleNameReq struct{}"
	}

	return strings.Join([]string{"RuleNameReq", string(data)}, " ")
}
