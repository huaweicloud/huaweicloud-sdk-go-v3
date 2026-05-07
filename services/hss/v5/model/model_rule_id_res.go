package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleIdRes **参数解释**： 规则ID **取值范围**： 字符长度1-36位
type RuleIdRes struct {
}

func (o RuleIdRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleIdRes struct{}"
	}

	return strings.Join([]string{"RuleIdRes", string(data)}, " ")
}
