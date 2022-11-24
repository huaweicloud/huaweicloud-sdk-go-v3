package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginOperApiInfo struct {

	// 绑定API的环境编码。
	EnvId string `json:"env_id"`

	// 绑定的API编码列表。
	ApiIds []string `json:"api_ids"`
}

func (o PluginOperApiInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginOperApiInfo struct{}"
	}

	return strings.Join([]string{"PluginOperApiInfo", string(data)}, " ")
}
