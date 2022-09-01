package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerSettingsDto struct {
	Configs *ContainerConfigsDto `json:"configs,omitempty" xml:"configs"`

	// 镜像存储地址
	ImageUrl string `json:"image_url" xml:"image_url"`

	// 环境变量
	Envs *interface{} `json:"envs,omitempty" xml:"envs"`

	// 卷配置
	Volumes *[]VolumeDto `json:"volumes,omitempty" xml:"volumes"`

	Resources *ResourceDto `json:"resources,omitempty" xml:"resources"`

	// 外挂设备配置
	ExtDevices *[]ExtDevice `json:"ext_devices,omitempty" xml:"ext_devices"`
}

func (o ContainerSettingsDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerSettingsDto struct{}"
	}

	return strings.Join([]string{"ContainerSettingsDto", string(data)}, " ")
}
