package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLlmGuardPolicyRulesResponse Response Object
type ListLlmGuardPolicyRulesResponse struct {

	// **参数解释：** Number of rules in the policy **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** Array of IpReputation rules **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]LlmRuleInfo `json:"items,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListLlmGuardPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLlmGuardPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListLlmGuardPolicyRulesResponse", string(data)}, " ")
}
