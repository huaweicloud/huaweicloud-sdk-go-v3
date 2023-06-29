package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopsRequest Request Object
type ListDesktopsRequest struct {

	// 桌面所属用户。
	UserName *string `json:"user_name,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 桌面IP地址。
	DesktopIp *string `json:"desktop_ip,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，取值范围0-1000，默认值1000。
	Limit *int32 `json:"limit,omitempty"`

	// 桌面池ID,多个桌面池ID用逗号隔开。
	PoolId *string `json:"pool_id,omitempty"`
}

func (o ListDesktopsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopsRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopsRequest", string(data)}, " ")
}
