package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateUserGroupRequest Request Object
type UpdateUserGroupRequest struct {

	// 桌面用户组ID。
	GroupId string `json:"group_id"`

	Body *EditUserGroupRequest `json:"body,omitempty"`
}

func (o UpdateUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateUserGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateUserGroupRequest", string(data)}, " ")
}
