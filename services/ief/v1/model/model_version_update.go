package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 应用模板版本配置
type VersionUpdate struct {
	// 环境变量

	Envs *[]Env `json:"envs,omitempty"`
	// 卷配置

	Volumes *[]Volumes `json:"volumes,omitempty"`

	Configs *VersionUpdateConfigs `json:"configs,omitempty"`

	Resources *VersionDetailResources `json:"resources,omitempty"`
	// 启动命令

	Command *[]string `json:"command,omitempty"`
	// 参数

	Args *[]string `json:"args,omitempty"`
	// 工作负载存活探针

	LivenessProbe *string `json:"liveness_probe,omitempty"`
	// 工作负载业务探针

	ReadinessProbe *string `json:"readiness_probe,omitempty"`
	// npu类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。

	NpuType *string `json:"npu_type,omitempty"`
}

func (o VersionUpdate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VersionUpdate struct{}"
	}

	return strings.Join([]string{"VersionUpdate", string(data)}, " ")
}
