package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用模板版本配置
type VersionDetail struct {

	// 应用版本号
	Version string `json:"version" xml:"version"`

	// 镜像存储地址
	ImageUrl string `json:"image_url" xml:"image_url"`

	// 环境变量
	Envs *[]Env `json:"envs,omitempty" xml:"envs"`

	// 卷配置
	Volumes *[]Volumes `json:"volumes,omitempty" xml:"volumes"`

	Configs *AppConfigs `json:"configs,omitempty" xml:"configs"`

	Resources *Resources `json:"resources,omitempty" xml:"resources"`

	// 架构
	Arch *string `json:"arch,omitempty" xml:"arch"`

	// 启动命令
	Command *[]string `json:"command,omitempty" xml:"command"`

	// 参数
	Args *[]string `json:"args,omitempty" xml:"args"`

	LivenessProbe *ProbeDetail `json:"liveness_probe,omitempty" xml:"liveness_probe"`

	ReadinessProbe *ProbeDetail `json:"readiness_probe,omitempty" xml:"readiness_probe"`

	// NPU类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。
	NpuType *string `json:"npu_type,omitempty" xml:"npu_type"`
}

func (o VersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetail struct{}"
	}

	return strings.Join([]string{"VersionDetail", string(data)}, " ")
}
