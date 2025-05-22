package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FlavorVolumeNodeInfo **参数解释**： 规格磁盘信息。 **取值范围**： 不涉及。
type FlavorVolumeNodeInfo struct {

	// **参数解释**： 节点使用存储类型。 **取值范围**： 不涉及。
	VolumeType string `json:"volume_type"`

	// **参数解释**： 节点使用的磁盘数量。 **取值范围**： 不涉及。
	VolumeNum int32 `json:"volume_num"`

	// **参数解释**： 节点去除副本后的有效容量。 **取值范围**： 不涉及。
	Capacity int32 `json:"capacity"`

	// **参数解释**： 节点存的单盘容量。 **取值范围**： 不涉及。
	VolumeSize int32 `json:"volume_size"`
}

func (o FlavorVolumeNodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FlavorVolumeNodeInfo struct{}"
	}

	return strings.Join([]string{"FlavorVolumeNodeInfo", string(data)}, " ")
}
