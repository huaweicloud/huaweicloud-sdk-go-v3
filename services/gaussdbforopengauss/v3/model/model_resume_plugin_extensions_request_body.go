package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResumePluginExtensionsRequestBody struct {

	// 插件名称
	PluginName string `json:"plugin_name"`

	// 数据库列表
	DbList []string `json:"db_list"`

	// 拓展模块名称
	ExtensionName string `json:"extension_name"`

	// 扩展开关。on表示开启，off表示关闭。
	ExtensionAction string `json:"extension_action"`
}

func (o ResumePluginExtensionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResumePluginExtensionsRequestBody struct{}"
	}

	return strings.Join([]string{"ResumePluginExtensionsRequestBody", string(data)}, " ")
}
