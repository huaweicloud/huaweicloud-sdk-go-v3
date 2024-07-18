package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListVpnUserGroupsResponse Response Object
type ListVpnUserGroupsResponse struct {

	// 用户组列表信息
	UserGroups *[]VpnUserGroup `json:"user_groups,omitempty"`

	// 总数
	TotalCount *int32 `json:"total_count,omitempty"`

	PageInfo *PageInfo `json:"page_info,omitempty"`

	// 请求ID
	RequestId      *string `json:"request_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListVpnUserGroupsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListVpnUserGroupsResponse struct{}"
	}

	return strings.Join([]string{"ListVpnUserGroupsResponse", string(data)}, " ")
}
