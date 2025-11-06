package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCcPolicyRulesResponse Response Object
type ListCcPolicyRulesResponse struct {

	// Number of rules in the policy
	Total *int32 `json:"total,omitempty"`

	// Array of Cc rules
	Items          *[]CcrulesListInfo `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCcPolicyRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCcPolicyRulesResponse struct{}"
	}

	return strings.Join([]string{"ListCcPolicyRulesResponse", string(data)}, " ")
}
