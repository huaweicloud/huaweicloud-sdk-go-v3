package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ControlItem struct {

	// 桌面id。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名称。
	DesktopName *string `json:"desktop_name,omitempty"`

	// 桌面状态。
	DesktopStatus *string `json:"desktop_status,omitempty"`

	// 池id。
	PoolId *string `json:"pool_id,omitempty"`

	// 桌面已分配的用户信息列表。
	AttachUserInfos *[]AttachInstancesUserInfo `json:"attach_user_infos,omitempty"`
}

func (o ControlItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ControlItem struct{}"
	}

	return strings.Join([]string{"ControlItem", string(data)}, " ")
}
