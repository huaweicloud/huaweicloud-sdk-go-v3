package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteWhiteBlackIpRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeleteWhiteBlackIpRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWhiteBlackIpRuleResponse struct{}"
	}

	return strings.Join([]string{"DeleteWhiteBlackIpRuleResponse", string(data)}, " ")
}
