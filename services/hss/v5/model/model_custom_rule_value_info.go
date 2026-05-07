package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomRuleValueInfo struct {

	// **参数解释**： 规则类型 **约束限制**： 必填 **取值范围**： - black_hash：黑hash  **默认取值**： 不涉及
	RuleType string `json:"rule_type"`

	// **参数解释**： hash类型 **约束限制**： 必填 **取值范围**： - SHA-256：sha256sum - MD5：md5sum - SHA-1：sha1sum  **默认取值**： 不涉及
	HashType string `json:"hash_type"`

	// **参数解释**： 是否自动阻断告警 **约束限制**： 必填 **取值范围**： - 0：不自动阻断告警 - 1：自动阻断告警  **默认取值**： 不涉及
	AutoBlock int32 `json:"auto_block"`

	// **参数解释**： 规则集列表 **约束限制**: 必填 **取值范围**: 1-1000个规则值 **默认取值**: 不涉及
	RuleValues []string `json:"rule_values"`
}

func (o CustomRuleValueInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomRuleValueInfo struct{}"
	}

	return strings.Join([]string{"CustomRuleValueInfo", string(data)}, " ")
}
