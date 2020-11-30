/*
 * FunctionGraph
 *
 * API v2
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 函数挂载配置。
type MountConfig struct {
	MountUser *MountUser `json:"mount_user"`
	// 函数挂载列表。
	FuncMounts []FuncMount `json:"func_mounts"`
}

func (o MountConfig) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"MountConfig", string(data)}, " ")
}
