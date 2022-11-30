package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EnvParam struct {

	// 环境关联组件id
	ComponentId *string `json:"component_id,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 环境名称
	EnvName string `json:"env_name"`

	// 环境类型，取值：DEV、TEST、PRE、ONLINE
	EnvType string `json:"env_type"`

	// OS类型，取值：LINUX、WINDOWS
	OsType *string `json:"os_type,omitempty"`

	// 环境关联region
	Region string `json:"region"`

	// 注册类型，取值：API、SERVICE_DISCOVERY、CONSOLE，默认值：API
	RegisterType *string `json:"register_type,omitempty"`
}

func (o EnvParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvParam struct{}"
	}

	return strings.Join([]string{"EnvParam", string(data)}, " ")
}
