package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrateResourceSpec 跨资源池迁移节点时目标资源池中节点的配置。
type MigrateResourceSpec struct {

	// **参数解释**：资源规格名称，跨资源池迁移时该参数必传。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Flavor string `json:"flavor"`

	CreatingStep *CreatingStep `json:"creatingStep,omitempty"`

	// **参数解释**：资源迁移的目标节点池名称。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodePool *string `json:"nodePool,omitempty"`

	RootVolume *RootVolume `json:"rootVolume,omitempty"`

	// **参数解释**：目标节点池的数据盘盘信息，新建节点池时有效。 **约束限制**：不涉及。
	DataVolumes *[]DataVolumeItem `json:"dataVolumes,omitempty"`

	// **参数解释**：磁盘高级配置。存在自定义数据盘时必须指定对应的高级配置，新建节点池时有效。 **约束限制**：不涉及。
	VolumeGroupConfigs *[]VolumeGroupConfig `json:"volumeGroupConfigs,omitempty"`

	// **参数解释**：k8s标签，格式为key/value键值对，非特权池不能指定。新建节点池时有效。 **约束限制**：不涉及。
	Labels map[string]string `json:"labels,omitempty"`

	// **参数解释**：支持给创建出来的节点加taints来设置反亲和性，非特权池不能指定。新建节点池时有效。 **约束限制**：不涉及。
	Taints *[]Taints `json:"taints,omitempty"`

	// **参数解释**：资源标签。新建节点池时有效。 **约束限制**：不涉及。
	Tags *[]UserTags `json:"tags,omitempty"`

	Network *NodeNetwork `json:"network,omitempty"`

	ExtendParams *ResourceExtendParams `json:"extendParams,omitempty"`
}

func (o MigrateResourceSpec) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrateResourceSpec struct{}"
	}

	return strings.Join([]string{"MigrateResourceSpec", string(data)}, " ")
}
