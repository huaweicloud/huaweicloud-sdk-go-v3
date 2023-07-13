package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MountComponent struct {

	// 环境ID。
	EnvId *string `json:"env_id,omitempty"`

	// 环境名称。
	EnvName *string `json:"env_name,omitempty"`

	// 应用ID。
	AppId *string `json:"app_id,omitempty"`

	// 应用名称。
	AppName *string `json:"app_name,omitempty"`

	// 组件ID。
	ComponentId *string `json:"component_id,omitempty"`

	// 组件名称。
	ComponentName *string `json:"component_name,omitempty"`
}

func (o MountComponent) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MountComponent struct{}"
	}

	return strings.Join([]string{"MountComponent", string(data)}, " ")
}
