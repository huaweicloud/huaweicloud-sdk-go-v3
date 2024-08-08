package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListServerGroupsRequest Request Object
type ListServerGroupsRequest struct {

	// 查询的偏移量。
	Offset *int32 `json:"offset,omitempty"`

	// 查询的数量，值区间[1-100]。
	Limit *int32 `json:"limit,omitempty"`

	// 服务器组名称。
	ServerGroupName *string `json:"server_group_name,omitempty"`

	// 服务器组唯一标识。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	// 应用组类型： * `SESSION_DESKTOP_APP` - 会话桌面app * `COMMON_APP` - 普通app
	AppType *string `json:"app_type,omitempty"`

	// 查询tag字段中包含该值的服务器组。
	Tags *string `json:"tags,omitempty"`

	// 企业项目ID(字段为空或者0表示使用默认default企业项目)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 是否为备服务器组，不传默认查所有： true : 是备服务器组。 false: 主服务器组，默认。
	IsSecondaryServerGroup *string `json:"is_secondary_server_group,omitempty"`
}

func (o ListServerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListServerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListServerGroupsRequest", string(data)}, " ")
}
