package model

import (
	"encoding/json"

	"strings"
)

// 标签。
type ListInstancesRespTags struct {
	// 标签的键。

	Key *string `json:"key,omitempty"`
	// 标签的值。

	Value *string `json:"value,omitempty"`
}

func (o ListInstancesRespTags) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListInstancesRespTags struct{}"
	}

	return strings.Join([]string{"ListInstancesRespTags", string(data)}, " ")
}
