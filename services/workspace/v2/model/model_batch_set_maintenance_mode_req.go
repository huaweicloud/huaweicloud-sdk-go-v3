package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchSetMaintenanceModeReq 批量设置桌面管理员维护模式。
type BatchSetMaintenanceModeReq struct {

	// 需要设置维护模式的desktopId列表。
	DesktopIds []string `json:"desktop_ids"`

	// 进入或退出管理员维护模式 false:  退出维护模式 true: 维护模式
	InMaintenanceMode bool `json:"in_maintenance_mode"`
}

func (o BatchSetMaintenanceModeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchSetMaintenanceModeReq struct{}"
	}

	return strings.Join([]string{"BatchSetMaintenanceModeReq", string(data)}, " ")
}
