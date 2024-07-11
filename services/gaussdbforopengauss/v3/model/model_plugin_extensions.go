package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginExtensions struct {

	// 拓展名称
	ExtensionName *string `json:"extension_name,omitempty"`

	// 拓展状态。on表示开启，off表示关闭。
	Status *string `json:"status,omitempty"`
}

func (o PluginExtensions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginExtensions struct{}"
	}

	return strings.Join([]string{"PluginExtensions", string(data)}, " ")
}
