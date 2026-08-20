package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExtensionModuleProperties 扩展模块属性信息
type ExtensionModuleProperties struct {

	// 构建清单版本
	BuildManifestVersion *string `json:"build_manifestVersion,omitempty"`

	// 名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 任务uuid
	Uuid *string `json:"uuid,omitempty"`

	// 操作系统
	OperationSystem *string `json:"operationSystem,omitempty"`

	// 镜像来源
	ImageSource *string `json:"imageSource,omitempty"`

	// 镜像名
	Image *string `json:"image,omitempty"`

	// 图标路径
	Icon *string `json:"icon,omitempty"`

	// 环境变量，按region映射。键为变量类别(如registry/mirror)，值为region到配置命令的映射。
	EnvironmentVariables map[string]map[string]string `json:"environmentVariables,omitempty"`

	Execution *ExtensionExecution `json:"execution,omitempty"`

	// 用户可配置参数列表。
	Parameters *[]ExtensionParameter `json:"parameters,omitempty"`

	// 内部标签。
	Tags *[]string `json:"tags,omitempty"`
}

func (o ExtensionModuleProperties) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExtensionModuleProperties struct{}"
	}

	return strings.Join([]string{"ExtensionModuleProperties", string(data)}, " ")
}
