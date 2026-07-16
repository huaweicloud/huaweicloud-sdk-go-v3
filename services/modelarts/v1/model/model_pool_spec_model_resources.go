package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoolSpecModelResources struct {

	// **参数解释**：资源规格ID。 **取值范围**：不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释**：资源池中资源规格实例数量。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：资源规格的弹性资源量。物理池中该值和count必须一致。 **取值范围**：不涉及。
	MaxCount *int32 `json:"maxCount,omitempty"`

	// **参数解释**：资源池中期望创建的资源规格实例的az分布。
	Azs *[]PoolNodeAz `json:"azs,omitempty"`

	ExtendParams *PoolSpecModelExtendParams `json:"extendParams,omitempty"`

	Os *Os `json:"os,omitempty"`

	DataVolumes *PoolSpecModelDataVolumes `json:"dataVolumes,omitempty"`

	VolumeGroupConfigs *PoolSpecModelVolumeGroupConfigs `json:"volumeGroupConfigs,omitempty"`
}

func (o PoolSpecModelResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolSpecModelResources struct{}"
	}

	return strings.Join([]string{"PoolSpecModelResources", string(data)}, " ")
}
