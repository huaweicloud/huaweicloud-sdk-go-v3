package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListProjectProtectedRefsUserGroupsResponse Response Object
type ListProjectProtectedRefsUserGroupsResponse struct {

	// 成员组列表
	Body *[]ProtectedRefsUserGroupBasicDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListProjectProtectedRefsUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListProjectProtectedRefsUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListProjectProtectedRefsUserGroupsResponse", string(data)}, " ")
}
