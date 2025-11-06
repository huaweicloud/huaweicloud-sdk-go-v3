package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupAddableUserGroupsResponse Response Object
type ListGroupAddableUserGroupsResponse struct {

	// **参数解释：** 成员组列表。
	Body           *[]UserGroupDto `json:"body,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListGroupAddableUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupAddableUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupAddableUserGroupsResponse", string(data)}, " ")
}
