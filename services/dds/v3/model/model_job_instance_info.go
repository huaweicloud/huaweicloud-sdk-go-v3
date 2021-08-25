package model

import (
	"encoding/json"

	"strings"
)

type JobInstanceInfo struct {
	// 实例ID。

	Id string `json:"id"`
	// 实例名称。

	Name string `json:"name"`
}

func (o JobInstanceInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "JobInstanceInfo struct{}"
	}

	return strings.Join([]string{"JobInstanceInfo", string(data)}, " ")
}
