package model

import (
	"encoding/json"

	"strings"
)

// 批量删除参数
type DeleteIds struct {
	// 所有删除对象id的集合

	Ids []string `json:"ids"`
}

func (o DeleteIds) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteIds struct{}"
	}

	return strings.Join([]string{"DeleteIds", string(data)}, " ")
}
