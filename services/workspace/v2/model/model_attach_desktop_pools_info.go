package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AttachDesktopPoolsInfo 关联桌面池信息
type AttachDesktopPoolsInfo struct {

	// 桌面池id
	DesktopPoolId *string `json:"desktop_pool_id,omitempty"`

	// 桌面池名称
	DesktopPoolName *string `json:"desktop_pool_name,omitempty"`

	// 是否分配了桌面
	IsAttached *bool `json:"is_attached,omitempty"`
}

func (o AttachDesktopPoolsInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachDesktopPoolsInfo struct{}"
	}

	return strings.Join([]string{"AttachDesktopPoolsInfo", string(data)}, " ")
}
