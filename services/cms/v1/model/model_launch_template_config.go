package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 启动模板相关配置
type LaunchTemplateConfig struct {
	LaunchTemplate *LaunchTemplateInfo `json:"launch_template"`

	// VM模板的覆盖
	Overrides []OverrideInfo `json:"overrides"`
}

func (o LaunchTemplateConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LaunchTemplateConfig struct{}"
	}

	return strings.Join([]string{"LaunchTemplateConfig", string(data)}, " ")
}
