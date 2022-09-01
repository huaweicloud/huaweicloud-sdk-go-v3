package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartInstanceParam struct {

	// 插件列表
	PluginEnableList *[]string `json:"plugin_enable_list,omitempty" xml:"plugin_enable_list"`

	// 插件参数
	PluginVars map[string]string `json:"plugin_vars,omitempty" xml:"plugin_vars"`
}

func (o StartInstanceParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceParam struct{}"
	}

	return strings.Join([]string{"StartInstanceParam", string(data)}, " ")
}
