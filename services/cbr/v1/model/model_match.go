package model

import (
	"encoding/json"

	"strings"
)

//
type Match struct {
	// 键。 key取值范围为：\"resource_name\"，资源名称

	Key string `json:"key"`
	// 值。 最大长度255个字符。 key为\"resource_name\"时，value为空字符串时精确匹配，为非空字符串时模糊匹配。

	Value string `json:"value"`
}

func (o Match) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Match struct{}"
	}

	return strings.Join([]string{"Match", string(data)}, " ")
}
