package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeSpec 节点期望信息。
type NodeSpec struct {

	// **参数解释**：节点资源规格ID。 **取值范围**：不涉及。
	Flavor string `json:"flavor"`

	Os *Os `json:"os,omitempty"`

	HostNetwork *NodeNetwork `json:"hostNetwork,omitempty"`

	RootVolume *VolumeVo `json:"rootVolume,omitempty"`

	// **参数解释**：数据盘信息。
	DataVolumes *[]VolumeVo `json:"dataVolumes,omitempty"`

	ExtendParams *ResourceExtendParams `json:"extendParams,omitempty"`
}

func (o NodeSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeSpec struct{}"
	}

	return strings.Join([]string{"NodeSpec", string(data)}, " ")
}
