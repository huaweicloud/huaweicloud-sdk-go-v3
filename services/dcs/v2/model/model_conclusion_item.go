package model

import (
	"encoding/json"

	"strings"
)

// 诊断结论
type ConclusionItem struct {
	// 结论id

	Id int32 `json:"id"`
	// 结论参数

	Params map[string]string `json:"params,omitempty"`
}

func (o ConclusionItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ConclusionItem struct{}"
	}

	return strings.Join([]string{"ConclusionItem", string(data)}, " ")
}
