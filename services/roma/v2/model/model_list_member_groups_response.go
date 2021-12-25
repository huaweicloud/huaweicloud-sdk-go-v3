package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListMemberGroupsResponse struct {
	// 本次返回的列表长度

	Size int32 `json:"size"`
	// 满足条件的记录数

	Total int64 `json:"total"`
	// VPC通道后端服务器组列表

	MemberGroups   *[]MemberGroupInfo `json:"member_groups,omitempty"`
	HttpStatusCode int                `json:"-"`
}

func (o ListMemberGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMemberGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListMemberGroupsResponse", string(data)}, " ")
}
