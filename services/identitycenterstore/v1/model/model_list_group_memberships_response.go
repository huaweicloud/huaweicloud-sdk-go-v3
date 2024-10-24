package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMembershipsResponse Response Object
type ListGroupMembershipsResponse struct {

	// 满足查询条件的关联关系对象列表
	GroupMemberships *[]GroupMembershipDto `json:"group_memberships,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListGroupMembershipsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMembershipsResponse struct{}"
	}

	return strings.Join([]string{"ListGroupMembershipsResponse", string(data)}, " ")
}
