package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleHitCountObject struct {

	// **参数解释**： 规则ID **取值范围**： 不涉及
	RuleId *string `json:"rule_id,omitempty"`

	// **参数解释**： 规则击中次数，当acl规则被触发时，对应规则ID的击中次数会添加一次。 **取值范围**： 不涉及
	RuleHitCount *int32 `json:"rule_hit_count,omitempty"`
}

func (o RuleHitCountObject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleHitCountObject struct{}"
	}

	return strings.Join([]string{"RuleHitCountObject", string(data)}, " ")
}
