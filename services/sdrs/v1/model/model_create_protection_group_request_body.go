package model

import (
	"encoding/json"

	"strings"
)

// 创建保护组请求体
type CreateProtectionGroupRequestBody struct {
	ServerGroup *CreateProtectionGroupRequestParams `json:"server_group"`
}

func (o CreateProtectionGroupRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateProtectionGroupRequestBody struct{}"
	}

	return strings.Join([]string{"CreateProtectionGroupRequestBody", string(data)}, " ")
}
