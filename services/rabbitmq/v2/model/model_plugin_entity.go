package model

import (
	"encoding/json"

	"strings"
)

// 插件信息。
type PluginEntity struct {
	// 是否运行。

	Running *bool `json:"running,omitempty"`
	// 是否启用。

	Enable *bool `json:"enable,omitempty"`
	// 插件名称。

	Name *string `json:"name,omitempty"`
	// 插件版本。

	Version *string `json:"version,omitempty"`
}

func (o PluginEntity) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PluginEntity struct{}"
	}

	return strings.Join([]string{"PluginEntity", string(data)}, " ")
}
