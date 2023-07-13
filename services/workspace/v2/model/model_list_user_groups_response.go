package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserGroupsResponse Response Object
type ListUserGroupsResponse struct {

	// 满足条件的用户组总数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 用户组列表。
	UserGroups     *[]UserGroupInfo `json:"user_groups,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ListUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListUserGroupsResponse", string(data)}, " ")
}
