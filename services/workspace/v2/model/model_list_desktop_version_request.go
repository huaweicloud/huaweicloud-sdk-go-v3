package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDesktopVersionRequest Request Object
type ListDesktopVersionRequest struct {

	// 桌面agent版本号（精确匹配）。
	AgentVersion string `json:"agent_version"`

	// 桌面操作系统类型。
	OsType string `json:"os_type"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称（支持模糊匹配）。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 用户名（支持模糊匹配）。
	Username *string `json:"username,omitempty"`

	// 用于分页查询，查询的起始记录序号，从0开始。
	Offset *int32 `json:"offset,omitempty"`

	// 用于分页查询，每页数量，默认10，最大100。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ListDesktopVersionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDesktopVersionRequest struct{}"
	}

	return strings.Join([]string{"ListDesktopVersionRequest", string(data)}, " ")
}
