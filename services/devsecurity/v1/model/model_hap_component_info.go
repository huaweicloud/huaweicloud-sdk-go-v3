package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type HapComponentInfo struct {

	// 权限列表
	Permission *[]PermissionItem `json:"permission,omitempty"`

	// 鸿蒙Page列表，仅鸿蒙任务存在该组件
	Page *[]HapComponentItem `json:"page,omitempty"`

	// 鸿蒙service列表，仅鸿蒙任务存在该组件
	Service *[]HapComponentItem `json:"service,omitempty"`

	// 鸿蒙data息列表，仅鸿蒙任务存在该组件
	Data *[]HapComponentItem `json:"data,omitempty"`
}

func (o HapComponentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HapComponentInfo struct{}"
	}

	return strings.Join([]string{"HapComponentInfo", string(data)}, " ")
}
