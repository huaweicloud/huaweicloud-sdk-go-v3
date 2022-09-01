package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 插件信息。
type PluginEntity struct {

	// 是否运行。
	Running *bool `json:"running,omitempty" xml:"running"`

	// 是否启用。
	Enable *bool `json:"enable,omitempty" xml:"enable"`

	// 插件名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 插件版本。
	Version *string `json:"version,omitempty" xml:"version"`
}

func (o PluginEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginEntity struct{}"
	}

	return strings.Join([]string{"PluginEntity", string(data)}, " ")
}
