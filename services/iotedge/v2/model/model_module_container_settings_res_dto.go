package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ModuleContainerSettingsResDto struct {
	Configs *ContainerConfigsResDto `json:"configs,omitempty"`

	// 自定义环境变量
	CustomEnvs *interface{} `json:"custom_envs,omitempty"`

	// 域名解析配置集合
	ExtraHosts *[]DnsConfigDto `json:"extra_hosts,omitempty"`
}

func (o ModuleContainerSettingsResDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ModuleContainerSettingsResDto struct{}"
	}

	return strings.Join([]string{"ModuleContainerSettingsResDto", string(data)}, " ")
}
