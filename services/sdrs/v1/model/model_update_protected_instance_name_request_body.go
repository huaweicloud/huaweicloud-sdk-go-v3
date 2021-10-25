package model

import (
	"encoding/json"

	"strings"
)

// 更新保护实例名称请求体
type UpdateProtectedInstanceNameRequestBody struct {
	ProtectedInstance *UpdateProtectedInstanceNameRequestParams `json:"protected_instance"`
}

func (o UpdateProtectedInstanceNameRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProtectedInstanceNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProtectedInstanceNameRequestBody", string(data)}, " ")
}
