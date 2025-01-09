package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DesktopDimensionAttachInfo 桌面维度关联关系。
type DesktopDimensionAttachInfo struct {

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 计划分配用户数。
	UserNum *int32 `json:"user_num,omitempty"`

	// 计划分配用户名称，如果有多个用逗号隔开。
	UserName *string `json:"user_name,omitempty"`
}

func (o DesktopDimensionAttachInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DesktopDimensionAttachInfo struct{}"
	}

	return strings.Join([]string{"DesktopDimensionAttachInfo", string(data)}, " ")
}
