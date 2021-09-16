package model

import (
	"encoding/json"

	"strings"
)

// 更新边缘实例请求体
type UpdateInstanceRequestBody struct {
	Server *UpdateInstanceOption `json:"server"`
}

func (o UpdateInstanceRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequestBody", string(data)}, " ")
}
