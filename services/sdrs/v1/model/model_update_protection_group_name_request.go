package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type UpdateProtectionGroupNameRequest struct {
	// 保护组ID。

	ServerGroupId string `json:"server_group_id"`

	Body *UpdateProtectionGroupNameRequestBody `json:"body,omitempty"`
}

func (o UpdateProtectionGroupNameRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateProtectionGroupNameRequest struct{}"
	}

	return strings.Join([]string{"UpdateProtectionGroupNameRequest", string(data)}, " ")
}
