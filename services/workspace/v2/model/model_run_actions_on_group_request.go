package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunActionsOnGroupRequest Request Object
type RunActionsOnGroupRequest struct {

	// 桌面用户组ID。
	GroupId string `json:"group_id"`

	Body *ActionsOfUsersInGroupRequest `json:"body,omitempty"`
}

func (o RunActionsOnGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunActionsOnGroupRequest struct{}"
	}

	return strings.Join([]string{"RunActionsOnGroupRequest", string(data)}, " ")
}
