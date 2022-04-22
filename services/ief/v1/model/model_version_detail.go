package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用模板版本配置
type VersionDetail struct {

	// 应用版本号
	Version string `json:"version"`

	// 镜像存储地址
	ImageUrl string `json:"image_url"`

	// 环境变量
	Envs *[]Env `json:"envs,omitempty"`

	// 卷配置
	Volumes *[]Volumes `json:"volumes,omitempty"`

	Configs *VersionDetailConfigs `json:"configs,omitempty"`

	Resources *VersionDetailResources `json:"resources,omitempty"`

	// 架构
	Arch *string `json:"arch,omitempty"`

	// 启动命令
	Command *[]string `json:"command,omitempty"`

	// 参数
	Args *[]string `json:"args,omitempty"`

	LivenessProbe *AppVersionDetailLivenessProbe `json:"liveness_probe,omitempty"`

	ReadinessProbe *AppVersionDetailReadinessProbe `json:"readiness_probe,omitempty"`

	// npu类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。
	NpuType *string `json:"npu_type,omitempty"`
}

func (o VersionDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionDetail struct{}"
	}

	return strings.Join([]string{"VersionDetail", string(data)}, " ")
}
