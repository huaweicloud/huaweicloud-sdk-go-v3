package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListGroupMembershipsForMemberResponse Response Object
type ListGroupMembershipsForMemberResponse struct {

	// 满足查询条件的关联关系对象列表
	GroupMemberships *[]GroupMembershipItem `json:"group_memberships,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListGroupMembershipsForMemberResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListGroupMembershipsForMemberResponse struct{}"
	}

	return strings.Join([]string{"ListGroupMembershipsForMemberResponse", string(data)}, " ")
}
