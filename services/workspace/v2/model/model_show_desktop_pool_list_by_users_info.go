package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDesktopPoolListByUsersInfo 每个用户所关联的桌面池列表信息
type ShowDesktopPoolListByUsersInfo struct {

	// 用户id
	UserId *string `json:"user_id,omitempty"`

	// 桌面池数
	DesktopPoolCount *int32 `json:"desktop_pool_count,omitempty"`

	// 运行状态统计
	DesktopPools *[]AttachDesktopPoolsInfo `json:"desktop_pools,omitempty"`
}

func (o ShowDesktopPoolListByUsersInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDesktopPoolListByUsersInfo struct{}"
	}

	return strings.Join([]string{"ShowDesktopPoolListByUsersInfo", string(data)}, " ")
}
