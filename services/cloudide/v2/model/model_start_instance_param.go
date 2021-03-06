package model

import (
	"encoding/json"

	"strings"
)

type StartInstanceParam struct {
	// 插件列表

	PluginEnableList *[]string `json:"plugin_enable_list,omitempty"`
	// 插件参数

	PluginVars map[string]string `json:"plugin_vars,omitempty"`
}

func (o StartInstanceParam) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "StartInstanceParam struct{}"
	}

	return strings.Join([]string{"StartInstanceParam", string(data)}, " ")
}
