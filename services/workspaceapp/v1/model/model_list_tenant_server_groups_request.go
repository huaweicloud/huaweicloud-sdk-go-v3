package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTenantServerGroupsRequest Request Object
type ListTenantServerGroupsRequest struct {

	// 查询的偏移量，默认值0。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]，默认值10。
	Limit *int32 `json:"limit,omitempty"`

	// 服务器组名称。
	SeverGroupName *string `json:"sever_group_name,omitempty"`

	// 应用组类型： * `SESSION_DESKTOP_APP` - 会话桌面app * `COMMON_APP` - 普通app
	AppType *string `json:"app_type,omitempty"`

	// 是否为备服务器组，不传默认查所有： true : 是备服务器组。 false: 主服务器组，默认。
	IsSecondaryServerGroup *string `json:"is_secondary_server_group,omitempty"`
}

func (o ListTenantServerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTenantServerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListTenantServerGroupsRequest", string(data)}, " ")
}
