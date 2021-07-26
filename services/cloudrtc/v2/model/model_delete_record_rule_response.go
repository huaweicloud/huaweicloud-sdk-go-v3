package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRecordRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteRecordRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteRecordRuleResponse", string(data)}, " ")
}
