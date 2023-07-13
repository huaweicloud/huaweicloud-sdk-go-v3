package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnusedDesktopInfo 未使用的桌面信息。
type UnusedDesktopInfo struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	ComputeName *string `json:"compute_name,omitempty"`

	// 桌面创建时间。
	CreateTime *string `json:"create_time,omitempty"`

	// 最近一次断连时间。
	DisconnectTime *string `json:"disconnect_time,omitempty"`
}

func (o UnusedDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnusedDesktopInfo struct{}"
	}

	return strings.Join([]string{"UnusedDesktopInfo", string(data)}, " ")
}
