package model

import (
	"encoding/json"

	"strings"
)

type Match struct {
	// 键。当前只有resource_name可用，表示集群的名称，后续扩展。

	Key string `json:"key"`
	// 值。每个值最大长度64个unicode字符。

	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
