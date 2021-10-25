package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowProtectionGroupRequest struct {
	// 保护组的ID。

	ServerGroupId string `json:"server_group_id"`
}

func (o ShowProtectionGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"ShowProtectionGroupRequest", string(data)}, " ")
}
