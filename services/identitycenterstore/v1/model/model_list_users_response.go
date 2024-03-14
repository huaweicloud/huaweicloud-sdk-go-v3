package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersResponse Response Object
type ListUsersResponse struct {
	PageInfo *PageInfoDto `json:"page_info,omitempty"`

	// 身份源中的用户列表
	Users          *[]UserDto `json:"users,omitempty"`
	HttpStatusCode int        `json:"-"`
}

func (o ListUsersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersResponse struct{}"
	}

	return strings.Join([]string{"ListUsersResponse", string(data)}, " ")
}
