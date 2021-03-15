package model

import (
	"encoding/json"

	"strings"
)

type ProxyReadonlyInstances struct {
	// 只读实例ID。
	Id string `json:"id"`
	// 只读实例权重。
	Weight int32 `json:"weight"`
}

func (o ProxyReadonlyInstances) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProxyReadonlyInstances struct{}"
	}

	return strings.Join([]string{"ProxyReadonlyInstances", string(data)}, " ")
}
