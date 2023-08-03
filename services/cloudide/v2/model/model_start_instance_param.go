package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type StartInstanceParam struct {

	// 插件列表
	PluginEnableList *[]string `json:"plugin_enable_list,omitempty"`

	// 插件参数，请注意敏感信息保护，若涉及敏感信息，请自行加密
	PluginVars map[string]string `json:"plugin_vars,omitempty"`
}

func (o StartInstanceParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartInstanceParam struct{}"
	}

	return strings.Join([]string{"StartInstanceParam", string(data)}, " ")
}
