package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersOfGroupResponse Response Object
type ListUsersOfGroupResponse struct {

	// 组关联的用户列表。
	Users *[]UserInGroup `json:"users,omitempty"`

	// 用户列表中用户总数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListUsersOfGroupResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersOfGroupResponse struct{}"
	}

	return strings.Join([]string{"ListUsersOfGroupResponse", string(data)}, " ")
}
