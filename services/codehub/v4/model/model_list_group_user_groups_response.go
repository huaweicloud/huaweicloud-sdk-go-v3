package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupUserGroupsResponse Response Object
type ListGroupUserGroupsResponse struct {

	// 成员组列表
	Body           *[]ProjectUserGroupDto `json:"body,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ListGroupUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupUserGroupsResponse", string(data)}, " ")
}
