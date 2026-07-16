package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DockerLvmConfig 节点容器磁盘配置项信息。
type DockerLvmConfig struct {

	// **参数解释**： 资源池节点Docker盘占数据盘的百分比。例如Docker盘大小占用数据盘20%，该参数值为20。 **取值范围**： 不涉及。
	DockerThinPool int32 `json:"dockerThinPool"`

	// **参数解释**： 资源池节点上kubelet占数据盘的百分比。例如Docker盘大小占用数据盘20%，该参数值为20。 **取值范围**： 不涉及。
	KubernetesLV int32 `json:"kubernetesLV"`

	// **参数解释**： 磁盘类型。 **取值范围**： 可选值如下： - evs：云硬盘 - ssd：本地SSD硬盘
	DockerDiskType string `json:"dockerDiskType"`

	// **参数解释**： LVM写入模式。 **取值范围**： 可选值如下： - striped：条带模式，使用多块磁盘组成条带模式，能够提升磁盘性能 - linear：线性模式
	LvType *string `json:"lvType,omitempty"`
}

func (o DockerLvmConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DockerLvmConfig struct{}"
	}

	return strings.Join([]string{"DockerLvmConfig", string(data)}, " ")
}
