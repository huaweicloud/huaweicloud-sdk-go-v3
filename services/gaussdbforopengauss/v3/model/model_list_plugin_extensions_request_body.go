package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPluginExtensionsRequestBody struct {

	// 数据库名称
	DbName string `json:"db_name"`

	// 插件名称
	PluginName string `json:"plugin_name"`
}

func (o ListPluginExtensionsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginExtensionsRequestBody struct{}"
	}

	return strings.Join([]string{"ListPluginExtensionsRequestBody", string(data)}, " ")
}
