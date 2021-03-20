package model

import (
	"encoding/json"

	"strings"
)

//
type RuleRef struct {
	// 规则ID。

	Id string `json:"id"`
}

func (o RuleRef) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RuleRef struct{}"
	}

	return strings.Join([]string{"RuleRef", string(data)}, " ")
}
