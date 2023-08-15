package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListCcRulesResponse Response Object
type ListCcRulesResponse struct {

	// Number of rules in the policy
	Total *int32 `json:"total,omitempty"`

	// Array of Cc rules
	Items          *[]CcrulesListInfo `json:"items,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListCcRulesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListCcRulesResponse struct{}"
	}

	return strings.Join([]string{"ListCcRulesResponse", string(data)}, " ")
}
