package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartProtectionGroupRequest struct {
	// 保护组的ID。

	ServerGroupId string `json:"server_group_id"`

	Body *StartProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o StartProtectionGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"StartProtectionGroupRequest", string(data)}, " ")
}
