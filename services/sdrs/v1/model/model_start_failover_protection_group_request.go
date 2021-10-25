package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type StartFailoverProtectionGroupRequest struct {
	// 保护组的ID。

	ServerGroupId string `json:"server_group_id"`

	Body *FailoverProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o StartFailoverProtectionGroupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartFailoverProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"StartFailoverProtectionGroupRequest", string(data)}, " ")
}
