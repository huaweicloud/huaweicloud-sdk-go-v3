package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Cause struct {

	// **参数解释**：策略名称。 **取值范围**：不涉及。
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**：策略条件。
	Condition *[]Conditions `json:"condition,omitempty"`
}

func (o Cause) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Cause struct{}"
	}

	return strings.Join([]string{"Cause", string(data)}, " ")
}
