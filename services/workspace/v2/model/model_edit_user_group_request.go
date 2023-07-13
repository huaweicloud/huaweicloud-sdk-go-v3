package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditUserGroupRequest struct {

	// 用户组描述。
	Description *string `json:"description,omitempty"`

	// 用户组名，注意AD用户组不支持改名。
	GroupName *string `json:"group_name,omitempty"`
}

func (o EditUserGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditUserGroupRequest struct{}"
	}

	return strings.Join([]string{"EditUserGroupRequest", string(data)}, " ")
}
