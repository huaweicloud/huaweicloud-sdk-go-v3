package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddonInstanceStatus 插件实例的状态
type AddonInstanceStatus struct {

	// 状态信息
	Status *string `json:"status,omitempty"`

	// 变化原因信息
	Reason *string `json:"Reason,omitempty"`

	// 变化详细信息
	Message *string `json:"message,omitempty"`

	// 目标版本信息
	TargetVersions *[]string `json:"targetVersions,omitempty"`

	CurrentVersion *AddonVersion `json:"currentVersion,omitempty"`
}

func (o AddonInstanceStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddonInstanceStatus struct{}"
	}

	return strings.Join([]string{"AddonInstanceStatus", string(data)}, " ")
}
