package model

import (
	"encoding/json"

	"strings"
)

type Nameserver struct {
	// 主机名。

	Hostname *string `json:"hostname,omitempty"`
	// 优先级。

	Priority *int32 `json:"priority,omitempty"`
}

func (o Nameserver) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Nameserver struct{}"
	}

	return strings.Join([]string{"Nameserver", string(data)}, " ")
}
