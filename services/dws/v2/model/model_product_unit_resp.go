package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProductUnitResp **参数解释**： 快照规格信息单个详情。 **取值范围**： 不涉及。
type ProductUnitResp struct {

	// **参数解释**： 规格ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规格编码。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 规格默认容量。 **取值范围**： 不涉及。
	DefaultCapacity *string `json:"default_capacity,omitempty"`

	// **参数解释**： 规格类型。 **取值范围**： 不涉及。
	Classify *string `json:"classify,omitempty"`

	// **参数解释**： 规格适用场景。 **取值范围**： 不涉及。
	Scenario *string `json:"scenario,omitempty"`

	// **参数解释**： 规格版本信息。 **取值范围**： v1.0：一代规格。 v2.0：二代规格。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 规格状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 扩展信息。 **取值范围**： 不涉及。
	Attribute *[]ProductExtendResp `json:"attribute,omitempty"`

	// **参数解释**： 规格使用副本数量。 **取值范围**： 不涉及。
	Duplicate *int32 `json:"duplicate,omitempty"`

	// **参数解释**： 默认节点数量。 **取值范围**： 不涉及。
	DefaultNode *int32 `json:"default_node,omitempty"`

	// **参数解释**： 最小节点数。 **取值范围**： 不涉及。
	MinNode *int32 `json:"min_node,omitempty"`

	// **参数解释**： 最大节点数。 **取值范围**： 不涉及。
	MaxNode *int32 `json:"max_node,omitempty"`

	// **参数解释**： 版本信息。 **取值范围**： 不涉及。
	ProductVersionList *[]ProductVersionResp `json:"product_version_list,omitempty"`

	// **参数解释**： 底层规格ID。有别于id字段，一般不会用到。 **取值范围**： 不涉及。
	FlavorId *string `json:"flavor_id,omitempty"`

	// **参数解释**： 规格的底层规格编码。 **取值范围**： 不涉及。
	FlavorCode *string `json:"flavor_code,omitempty"`

	// **参数解释**： 规格的磁盘数。 **取值范围**： 不涉及。
	VolumeNum *int32 `json:"volume_num,omitempty"`

	VolumeUsed *ProductVolumeUsedResp `json:"volume_used,omitempty"`
}

func (o ProductUnitResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProductUnitResp struct{}"
	}

	return strings.Join([]string{"ProductUnitResp", string(data)}, " ")
}
