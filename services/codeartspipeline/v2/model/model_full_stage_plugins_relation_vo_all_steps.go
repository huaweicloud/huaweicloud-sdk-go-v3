package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoAllSteps struct {

	// 插件名
	PluginName *string `json:"plugin_name,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 版本
	Version *string `json:"version,omitempty"`
}

func (o FullStagePluginsRelationVoAllSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoAllSteps struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoAllSteps", string(data)}, " ")
}
