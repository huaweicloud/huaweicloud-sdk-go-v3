package model

import (
	"encoding/json"

	"strings"
)

type Target struct {
	// 实例ID。

	InstanceId string `json:"instance_id"`
}

func (o Target) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Target struct{}"
	}

	return strings.Join([]string{"Target", string(data)}, " ")
}
