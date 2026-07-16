package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolResource 资源池资源规格资源量及节点池配置数据模型。
type PoolResource struct {

	// **参数解释**：资源规格名称，比如：modelarts.vm.gpu.t4u8。 **取值范围**：不涉及。
	Flavor string `json:"flavor"`

	// **参数解释**：规格保障使用量。 **取值范围**：不涉及。
	Count int32 `json:"count"`

	// **参数解释**：资源规格的弹性使用量，物理池该值和count相同[，逻辑池该值大于等于count](tag:hcs,hcso)。 **取值范围**：不涉及。
	MaxCount int32 `json:"maxCount"`

	// **参数解释**：资源池中节点的AZ信息。
	Azs *[]PoolNodeAz `json:"azs,omitempty"`

	// **参数解释**：节点池名称。比如：nodePool-1。 **取值范围**：不涉及。
	NodePool *string `json:"nodePool,omitempty"`

	// **参数解释**：支持给创建出来的节点加taints来设置反亲和性，非特权池不能指定。
	Taints *[]Taints `json:"taints,omitempty"`

	// **参数解释**：k8s标签，格式为key/value键值对。
	Labels map[string]string `json:"labels,omitempty"`

	// **参数解释**：资源标签，非特权池不能指定。
	Tags *[]UserTags `json:"tags,omitempty"`

	Network *NodeNetwork `json:"network,omitempty"`

	ExtendParams *ResourceExtendParams `json:"extendParams,omitempty"`

	CreatingStep *CreatingStep `json:"creatingStep,omitempty"`

	RootVolume *RootVolume `json:"rootVolume,omitempty"`

	// **参数解释**：自定义数据盘（云硬盘）列表信息。
	DataVolumes *[]DataVolumeItem `json:"dataVolumes,omitempty"`

	// **参数解释**：磁盘高级配置。存在自定义数据盘时必须指定对应的高级配置。
	VolumeGroupConfigs *[]VolumeGroupConfig `json:"volumeGroupConfigs,omitempty"`

	Os *Os `json:"os,omitempty"`
}

func (o PoolResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolResource struct{}"
	}

	return strings.Join([]string{"PoolResource", string(data)}, " ")
}
