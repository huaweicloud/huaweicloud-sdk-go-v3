package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartInstanceParam struct {
	// 插件列表

	PluginEnableList *[]string `json:"plugin_enable_list,omitempty"`
	// 插件参数

	PluginVars map[string]string `json:"plugin_vars,omitempty"`
}

func (o StartInstanceParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceParam struct{}"
	}

	return strings.Join([]string{"StartInstanceParam", string(data)}, " ")
}
