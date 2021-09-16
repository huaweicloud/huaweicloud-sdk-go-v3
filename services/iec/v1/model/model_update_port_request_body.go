package model

import (
	"encoding/json"

	"strings"
)

// 更新端口请求体
type UpdatePortRequestBody struct {
	Port *UpdatePortOption `json:"port"`
}

func (o UpdatePortRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePortRequestBody struct{}"
	}

	return strings.Join([]string{"UpdatePortRequestBody", string(data)}, " ")
}
