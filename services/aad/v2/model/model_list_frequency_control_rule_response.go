package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListFrequencyControlRuleResponse Response Object
type ListFrequencyControlRuleResponse struct {

	// total
	Total *int32 `json:"total,omitempty"`

	// items
	Items          *[]FrequencyControlRule `json:"items,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListFrequencyControlRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListFrequencyControlRuleResponse struct{}"
	}

	return strings.Join([]string{"ListFrequencyControlRuleResponse", string(data)}, " ")
}
