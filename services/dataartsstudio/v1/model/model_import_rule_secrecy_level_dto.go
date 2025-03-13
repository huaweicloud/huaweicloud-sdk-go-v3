package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ImportRuleSecrecyLevelDto struct {

	// 内置规则模板id。
	BuiltinRuleId *string `json:"builtin_rule_id,omitempty"`

	// 密级id，获取方法请参见[获取数据密级](ListSecuritySecrecyLevels.xml)。
	SecrecyLevel *string `json:"secrecy_level,omitempty"`
}

func (o ImportRuleSecrecyLevelDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportRuleSecrecyLevelDto struct{}"
	}

	return strings.Join([]string{"ImportRuleSecrecyLevelDto", string(data)}, " ")
}
