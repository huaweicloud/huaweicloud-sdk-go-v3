package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopProtectionGroupRequest struct {

	// 保护组的ID。
	ServerGroupId string `json:"server_group_id"`

	Body *StopProtectionGroupRequestBody `json:"body,omitempty"`
}

func (o StopProtectionGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopProtectionGroupRequest struct{}"
	}

	return strings.Join([]string{"StopProtectionGroupRequest", string(data)}, " ")
}
