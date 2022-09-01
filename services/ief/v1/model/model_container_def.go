package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ContainerDef struct {

	// 容器名称，只允许英文小写字母、数字、中划线，最大长度32， 英文小写字母或数字开头和结尾
	Name string `json:"name" xml:"name"`

	// 容器镜像URL
	ImageUrl string `json:"image_url" xml:"image_url"`

	// 容器启动参数，字符总长度最大为65536
	Args *[]string `json:"args,omitempty" xml:"args"`

	// 容器启动命令，字符总长度最大为65536。 command支持使用数组定义多条命令，但在IEF控制台界面只会显示第一条命令。
	Command *[]string `json:"command,omitempty" xml:"command"`

	Resources *DeploymentResources `json:"resources,omitempty" xml:"resources"`

	// 环境变量
	Envs *[]Env `json:"envs,omitempty" xml:"envs"`

	// 容器端口映射值
	Ports *[]HostContainerPortMapping `json:"ports,omitempty" xml:"ports"`

	// 是否启用特权容器,默认值false
	Privileged *bool `json:"privileged,omitempty" xml:"privileged"`

	ReadinessProbe *Probe `json:"readiness_probe,omitempty" xml:"readiness_probe"`

	LivenessProbe *Probe `json:"liveness_probe,omitempty" xml:"liveness_probe"`

	// 容器镜像版本
	Version *string `json:"version,omitempty" xml:"version"`

	// 卷配置
	Volumes *[]Volumes `json:"volumes,omitempty" xml:"volumes"`

	// NPU类型，支持D310类型和D910类型。 - D310表示D310类型。 - D910表示D910类型。 - 不填表示为D310类型。
	NpuType *string `json:"npu_type,omitempty" xml:"npu_type"`
}

func (o ContainerDef) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerDef struct{}"
	}

	return strings.Join([]string{"ContainerDef", string(data)}, " ")
}
