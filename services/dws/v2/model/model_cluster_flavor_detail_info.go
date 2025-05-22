package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ClusterFlavorDetailInfo **参数解释**： 集群规格详情。 **取值范围**： 不涉及。
type ClusterFlavorDetailInfo struct {

	// **参数解释**： 规格ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 规格编码。 **取值范围**： 不涉及。
	SpecName string `json:"spec_name"`

	// **参数解释**： 当前节点数量。 **取值范围**： 不涉及。
	CurrentNode int32 `json:"current_node"`

	// **参数解释**： 最小节点阈值。 **取值范围**： 不涉及。
	MinNode int32 `json:"min_node"`

	// **参数解释**： 最大节点阈值。 **取值范围**： 不涉及。
	MaxNode int32 `json:"max_node"`

	// **参数解释**： 规格类型。 **取值范围**： 不涉及。
	Classify string `json:"classify"`

	// **参数解释**： 数据仓库版本。 **取值范围**： 不涉及。
	DatastoreVersion string `json:"datastore_version"`

	// **参数解释**： 扩展信息。 **取值范围**： 不涉及。
	Attribute []FlavorAttributeInfo `json:"attribute"`

	VolumeNode *FlavorVolumeNodeInfo `json:"volume_node"`
}

func (o ClusterFlavorDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ClusterFlavorDetailInfo struct{}"
	}

	return strings.Join([]string{"ClusterFlavorDetailInfo", string(data)}, " ")
}
