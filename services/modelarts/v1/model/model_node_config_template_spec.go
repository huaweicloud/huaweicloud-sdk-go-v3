package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeConfigTemplateSpec 节点配置模板信息。
type NodeConfigTemplateSpec struct {

	// **参数解释**： 资源池节点上单容器的可用磁盘空间大小，单位G。 **取值范围**： 不涉及。
	DockerBaseSize int32 `json:"dockerBaseSize"`

	DockerLvmConfig *DockerLvmConfig `json:"dockerLvmConfig,omitempty"`

	// **参数解释**：该规格支持的Modelarts内置操作系统列表。
	OsList *[]AffinityOs `json:"osList,omitempty"`
}

func (o NodeConfigTemplateSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeConfigTemplateSpec struct{}"
	}

	return strings.Join([]string{"NodeConfigTemplateSpec", string(data)}, " ")
}
