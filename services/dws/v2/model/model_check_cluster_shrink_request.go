package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckClusterShrinkRequest Request Object
type CheckClusterShrinkRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 检查项，取值当前仅包含3种。 **约束限制**： 不涉及。 **取值范围**： - guc：检查当前guc参数是否满足缩容条件。 - schema：检查所有schema下有无影响缩容的表。 - disk：检查缩容后磁盘容量是否满足要求。  **默认取值**： 不涉及。
	CheckItem string `json:"check_item"`

	// **参数解释**： 待缩容节点数。 **约束限制**： 不涉及。 **取值范围**： 最小值为3，最大值为当前节点总数减3。 **默认取值**： 不涉及。
	ShrinkCount int32 `json:"shrink_count"`
}

func (o CheckClusterShrinkRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckClusterShrinkRequest struct{}"
	}

	return strings.Join([]string{"CheckClusterShrinkRequest", string(data)}, " ")
}
