package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginSpec 插件实例的具体信息。
type PluginSpec struct {
	Template *Template `json:"template"`
}

func (o PluginSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginSpec struct{}"
	}

	return strings.Join([]string{"PluginSpec", string(data)}, " ")
}
