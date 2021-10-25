package model

import (
	"encoding/json"

	"strings"
)

// 更新保护组名称请求体
type UpdateProtectionGroupNameRequestBody struct {
	ServerGroup *UpdateProtectionGroupNameRequestParams `json:"server_group"`
}

func (o UpdateProtectionGroupNameRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateProtectionGroupNameRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateProtectionGroupNameRequestBody", string(data)}, " ")
}
