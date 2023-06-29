package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteUserGroupRequest Request Object
type DeleteUserGroupRequest struct {

	// 桌面用户组ID。
	GroupId string `json:"group_id"`
}

func (o DeleteUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteUserGroupRequest struct{}"
	}

	return strings.Join([]string{"DeleteUserGroupRequest", string(data)}, " ")
}
