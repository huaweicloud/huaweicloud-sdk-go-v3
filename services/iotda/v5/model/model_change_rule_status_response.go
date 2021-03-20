package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ChangeRuleStatusResponse struct {
	// 规则的激活状态。 - active：激活。 - inactive：未激活。

	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeRuleStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"ChangeRuleStatusResponse", string(data)}, " ")
}
