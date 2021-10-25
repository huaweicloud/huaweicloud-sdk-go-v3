package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartReverseProtectionGroupRequest struct {
	// 保护组的ID。

	ServerGroupId string `json:"server_group_id"`

	Body *ReverseProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o StartReverseProtectionGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartReverseProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"StartReverseProtectionGroupRequest", string(data)}, " ")
}
