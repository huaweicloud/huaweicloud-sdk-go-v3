package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListPrivacyRuleResponse struct {
	// 规则条数

	Total *int32 `json:"total,omitempty"`
	// 规则

	Items          *[]PrivacyResponseBody `json:"items,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListPrivacyRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListPrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"ListPrivacyRuleResponse", string(data)}, " ")
}
