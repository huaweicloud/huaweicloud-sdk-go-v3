package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShrinkLogicalClusterRequest Request Object
type ShrinkLogicalClusterRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 必须是有效的dws集群ID。 **取值范围**： 36位UUID。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 逻辑集群id。  **约束限制**： 必须是有效的dws逻辑集群ID。  **取值范围**：  36位UUID。  **默认取值**：  不涉及。
	LogicalClusterId string `json:"logical_cluster_id"`

	Body *ShrinkLogicalClusterRequestBody `json:"body,omitempty"`
}

func (o ShrinkLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShrinkLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"ShrinkLogicalClusterRequest", string(data)}, " ")
}
