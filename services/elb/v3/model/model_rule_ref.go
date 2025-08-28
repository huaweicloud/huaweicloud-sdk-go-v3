package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RuleRef
type RuleRef struct {

	// **参数解释**：规则ID。  **取值范围**：不涉及
	Id string `json:"id"`
}

func (o RuleRef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleRef struct{}"
	}

	return strings.Join([]string{"RuleRef", string(data)}, " ")
}
