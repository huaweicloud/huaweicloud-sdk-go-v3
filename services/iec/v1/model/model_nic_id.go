package model

import (
	"encoding/json"

	"strings"
)

// 网卡ID信息。
type NicId struct {
	// 网卡ID。

	Id string `json:"id"`
}

func (o NicId) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NicId struct{}"
	}

	return strings.Join([]string{"NicId", string(data)}, " ")
}