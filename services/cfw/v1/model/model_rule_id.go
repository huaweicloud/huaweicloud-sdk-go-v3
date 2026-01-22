package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleId struct {

	// **参数解释**： 规则ID **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规则名称 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`
}

func (o RuleId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleId struct{}"
	}

	return strings.Join([]string{"RuleId", string(data)}, " ")
}
