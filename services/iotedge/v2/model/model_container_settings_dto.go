package model

import (
	"encoding/json"

	"strings"
)

type ContainerSettingsDto struct {
	Configs *ContainerConfigsDto `json:"configs,omitempty"`
	// 镜像存储地址

	ImageUrl string `json:"image_url"`
	// 环境变量

	Envs *interface{} `json:"envs,omitempty"`
	// 卷配置

	Volumes *[]VolumeDto `json:"volumes,omitempty"`

	Resources *ResourceDto `json:"resources,omitempty"`
}

func (o ContainerSettingsDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ContainerSettingsDto struct{}"
	}

	return strings.Join([]string{"ContainerSettingsDto", string(data)}, " ")
}
