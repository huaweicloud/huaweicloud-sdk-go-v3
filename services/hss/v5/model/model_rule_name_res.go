package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleNameRes **参数解释**： 规则名称 **取值范围**： 字符长度1-64位
type RuleNameRes struct {
}

func (o RuleNameRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleNameRes struct{}"
	}

	return strings.Join([]string{"RuleNameRes", string(data)}, " ")
}
