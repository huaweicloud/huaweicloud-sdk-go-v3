package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VulAffectContainerAppPath 受漏洞影响容器的软件路径
type VulAffectContainerAppPath struct {

	// 处理状态
	Status *string `json:"status,omitempty"`

	// 软件版本
	AppVersion *string `json:"app_version,omitempty"`

	// 软件路径
	AppPath *string `json:"app_path,omitempty"`
}

func (o VulAffectContainerAppPath) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VulAffectContainerAppPath struct{}"
	}

	return strings.Join([]string{"VulAffectContainerAppPath", string(data)}, " ")
}
