package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ListIgnoreRuleResponse struct {
	// 该策略下规则数量

	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListIgnoreRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListIgnoreRuleResponse struct{}"
	}

	return strings.Join([]string{"ListIgnoreRuleResponse", string(data)}, " ")
}
