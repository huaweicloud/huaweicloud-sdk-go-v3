package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlterAutoVolumeExpandConfig **参数解释**： 修改自动磁盘扩容配置 **约束限制**： 不涉及。
type AlterAutoVolumeExpandConfig struct {

	// **参数解释**： 是否开启磁盘自动扩容。 **约束限制**： 不涉及。 **取值范围**： - true：开启磁盘自动扩容。 - false：关闭磁盘自动扩容。 **默认取值**： 不涉及。
	AutoVolumeExpandEnable bool `json:"auto_volume_expand_enable"`

	// **参数解释**： 触发磁盘自动扩容的阈值。 **约束限制**： 不涉及。 **取值范围**： 20%-80%。 **默认取值**： 不涉及。
	ExpandThreshold *int32 `json:"expand_threshold,omitempty"`

	// **参数解释**： 磁盘自动扩容的上限值。 **约束限制**： - 能被100整除。 -大于当前实例磁盘容量，且小于节点数*30000。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxVolumeSize *int32 `json:"max_volume_size,omitempty"`

	// **参数解释**： 每次磁盘自动扩容的比例。 **约束限制**： 不涉及。 **取值范围**： 10%-100%。 **默认取值**： 不涉及。
	ExpandIncrement *int32 `json:"expand_increment,omitempty"`
}

func (o AlterAutoVolumeExpandConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlterAutoVolumeExpandConfig struct{}"
	}

	return strings.Join([]string{"AlterAutoVolumeExpandConfig", string(data)}, " ")
}
