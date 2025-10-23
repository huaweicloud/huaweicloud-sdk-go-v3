package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupProtectedRefsUserGroupsResponse Response Object
type ListGroupProtectedRefsUserGroupsResponse struct {

	// 成员组列表
	Body *[]ProtectedRefsUserGroupBasicDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListGroupProtectedRefsUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupProtectedRefsUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupProtectedRefsUserGroupsResponse", string(data)}, " ")
}
