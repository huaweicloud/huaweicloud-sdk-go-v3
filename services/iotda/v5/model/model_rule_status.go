package model

import (
	"encoding/json"

	"strings"
)

// 规则状态
type RuleStatus struct {
	// 规则的激活状态。 - active：激活。 - inactive：未激活。

	Status string `json:"status"`
}

func (o RuleStatus) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RuleStatus struct{}"
	}

	return strings.Join([]string{"RuleStatus", string(data)}, " ")
}
