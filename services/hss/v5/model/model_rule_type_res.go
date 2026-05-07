package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleTypeRes **参数解释**： 规则类型 **取值范围**： - black_hash：黑hash
type RuleTypeRes struct {
}

func (o RuleTypeRes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleTypeRes struct{}"
	}

	return strings.Join([]string{"RuleTypeRes", string(data)}, " ")
}
