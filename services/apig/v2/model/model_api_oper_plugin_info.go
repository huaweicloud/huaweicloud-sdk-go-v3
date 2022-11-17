package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ApiOperPluginInfo struct {

	// 绑定API的环境编码。
	EnvId string `json:"env_id"`

	// 绑定的插件编码列表。
	PluginIds []string `json:"plugin_ids"`
}

func (o ApiOperPluginInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ApiOperPluginInfo struct{}"
	}

	return strings.Join([]string{"ApiOperPluginInfo", string(data)}, " ")
}
