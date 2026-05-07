package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleStatusRes **参数解释**： 规则状态 **取值范围**: - 0：停用 - 1：启用
type RuleStatusRes struct {
}

func (o RuleStatusRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleStatusRes struct{}"
	}

	return strings.Join([]string{"RuleStatusRes", string(data)}, " ")
}
