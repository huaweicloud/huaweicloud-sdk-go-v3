package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRepositoryProtectedRefsUserGroupsResponse Response Object
type ListRepositoryProtectedRefsUserGroupsResponse struct {

	// 成员组列表
	Body *[]ProtectedRefsUserGroupBasicDto `json:"body,omitempty"`

	XTotal         *string `json:"X-Total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListRepositoryProtectedRefsUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRepositoryProtectedRefsUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListRepositoryProtectedRefsUserGroupsResponse", string(data)}, " ")
}
