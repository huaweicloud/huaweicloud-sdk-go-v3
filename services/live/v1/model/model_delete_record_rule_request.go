package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRecordRuleRequest struct {
	// 规则ID

	Id string `json:"id"`
}

func (o DeleteRecordRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordRuleRequest", string(data)}, " ")
}
