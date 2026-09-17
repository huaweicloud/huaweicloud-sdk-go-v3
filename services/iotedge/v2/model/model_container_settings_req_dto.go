package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerSettingsReqDto struct {
	Configs *ContainerConfigsReqDto `json:"configs,omitempty"`

	// 自定义环境变量
	CustomEnvs *interface{} `json:"custom_envs,omitempty"`

	// 域名解析配置集合。示例：[{\"hostname\":\"endpoint\",\"ip\":\"127.0.0.1\"}]
	ExtraHosts *interface{} `json:"extra_hosts,omitempty"`
}

func (o ContainerSettingsReqDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerSettingsReqDto struct{}"
	}

	return strings.Join([]string{"ContainerSettingsReqDto", string(data)}, " ")
}
