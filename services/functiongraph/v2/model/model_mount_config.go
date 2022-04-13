package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 函数挂载配置。
type MountConfig struct {
	MountUser *MountUser `json:"mount_user"`
	// 函数挂载列表。

	FuncMounts []FuncMount `json:"func_mounts"`
}

func (o MountConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MountConfig struct{}"
	}

	return strings.Join([]string{"MountConfig", string(data)}, " ")
}
