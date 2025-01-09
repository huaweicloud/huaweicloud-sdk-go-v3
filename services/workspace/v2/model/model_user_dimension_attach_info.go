package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UserDimensionAttachInfo 用户维度关联关系。
type UserDimensionAttachInfo struct {

	// 用户名称。
	UserName *string `json:"user_name,omitempty"`

	// 桌面id。
	UserId *string `json:"user_id,omitempty"`

	// 计划分配桌面数。
	DesktopNum *int32 `json:"desktop_num,omitempty"`

	// 计划分配桌面名称，如果有多个用逗号隔开。
	DesktopName *string `json:"desktop_name,omitempty"`
}

func (o UserDimensionAttachInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UserDimensionAttachInfo struct{}"
	}

	return strings.Join([]string{"UserDimensionAttachInfo", string(data)}, " ")
}
