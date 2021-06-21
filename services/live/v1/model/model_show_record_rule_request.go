package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRecordRuleRequest struct {
	// 规则ID

	Id string `json:"id"`
}

func (o ShowRecordRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowRecordRuleRequest", string(data)}, " ")
}
