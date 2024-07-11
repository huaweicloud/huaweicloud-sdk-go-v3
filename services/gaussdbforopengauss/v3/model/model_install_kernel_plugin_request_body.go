package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstallKernelPluginRequestBody struct {

	// 插件名称
	PluginName string `json:"plugin_name"`

	// 插件安装包链接
	Url string `json:"url"`

	// 插件安装包的sha256值
	Sha256 string `json:"sha_256"`
}

func (o InstallKernelPluginRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallKernelPluginRequestBody struct{}"
	}

	return strings.Join([]string{"InstallKernelPluginRequestBody", string(data)}, " ")
}
