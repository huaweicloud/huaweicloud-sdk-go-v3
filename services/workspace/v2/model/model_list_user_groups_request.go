package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUserGroupsRequest Request Object
type ListUserGroupsRequest struct {

	// 用于分页查询，返回用户组数量限制。如果不指定或为0，则使用默认值100，从1开始，最大100。
	Limit *string `json:"limit,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始，默认值0，必须与limit同时使用。
	Offset *string `json:"offset,omitempty"`

	// 用来匹配用户组的搜索关键字。例如根据组名模糊查询。
	Keyword *string `json:"keyword,omitempty"`

	// 根据用户组的域名进行过滤。
	Domain *string `json:"domain,omitempty"`

	// 用户类型。 * AD： AD域用户 * LOCAL： 本地liteAs用户 * UOS： UOS域用户
	PlatformType *[]string `json:"platform_type,omitempty"`
}

func (o ListUserGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUserGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListUserGroupsRequest", string(data)}, " ")
}
