package model

import (
	"encoding/json"

	"strings"
)

type InstanceRestartRequsetBody struct {
	// 在线调试时必填。

	Restart *interface{} `json:"restart,omitempty"`
}

func (o InstanceRestartRequsetBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "InstanceRestartRequsetBody struct{}"
	}

	return strings.Join([]string{"InstanceRestartRequsetBody", string(data)}, " ")
}
