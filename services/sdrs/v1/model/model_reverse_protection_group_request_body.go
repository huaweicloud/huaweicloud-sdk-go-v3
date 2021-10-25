package model

import (
	"encoding/json"

	"strings"
)

// 保护组切换请求体
type ReverseProtectionGroupRequestBody struct {
	ReverseServerGroup *ReverseProtectionGroupRequestParams `json:"reverse-server-group"`
}

func (o ReverseProtectionGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ReverseProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"ReverseProtectionGroupRequestBody", string(data)}, " ")
}
