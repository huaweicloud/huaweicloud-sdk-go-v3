package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ComponentConfig struct {

	// 组件名称
	ComponentName string `json:"component_name"`

	// 组件配置项列表
	Configs *[]Config `json:"configs,omitempty"`
}

func (o ComponentConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComponentConfig struct{}"
	}

	return strings.Join([]string{"ComponentConfig", string(data)}, " ")
}
