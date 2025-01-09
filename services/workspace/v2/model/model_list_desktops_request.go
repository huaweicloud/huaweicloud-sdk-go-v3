package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsRequest Request Object
type ListDesktopsRequest struct {

	// 桌面所属用户。
	UserName *[]string `json:"user_name,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面IP地址。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 桌面的sid列表，一次只能查询20个sid。
	Sids *[]string `json:"sids,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-1000，默认值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 用于筛选指定站点下的桌面列表
	SiteId *string `json:"site_id,omitempty"`

	// 桌面池ID,多个桌面池ID用逗号隔开。
	PoolId *string `json:"pool_id,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 桌面类型，为空时查所有桌面。查询多个类型时用,隔开。 - DEDICATED：普通桌面，包括专享桌面、专属桌面等。 - SHARED: 多用户共享桌面。
	DesktopType *string `json:"desktop_type,omitempty"`

	// 是否为协同桌面
	IsShareDesktop *bool `json:"is_share_desktop,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 桌面的运行状态。
	Status *string `json:"status,omitempty"`

	// 桌面id，当前最多支持100个桌面id进行查询。
	DesktopId *[]string `json:"desktop_id,omitempty"`

	// 桌面的标签。样例： - key1=value1。 - key1=value1，key2=value2。
	Tag *string `json:"tag,omitempty"`
}

func (o ListDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsRequest", string(data)}, " ")
}
