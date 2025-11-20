package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PartitionSpec **参数解释**： 分区的配置信息 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
type PartitionSpec struct {
	HostNetwork *PartitionSpecHostNetwork `json:"hostNetwork,omitempty"`

	// **参数解释**： 分区容器子网 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	ContainerNetwork *[]PartitionSpecContainerNetwork `json:"containerNetwork,omitempty"`

	// **参数解释**： 群组，IES边缘场景为可用区ID，中心区统一为“center”。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	PublicBorderGroup *string `json:"publicBorderGroup,omitempty"`

	// **参数解释**： 可用区分类 **约束限制**： 不涉及 **取值范围**： - Default: 中心云可用区 - IES: 专属云可用区 - HomeZone: 智能边缘云可用区  **默认取值**： 不涉及
	Category *string `json:"category,omitempty"`
}

func (o PartitionSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PartitionSpec struct{}"
	}

	return strings.Join([]string{"PartitionSpec", string(data)}, " ")
}
