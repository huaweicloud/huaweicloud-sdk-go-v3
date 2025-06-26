package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestartLogicalClusterRequest Request Object
type RestartLogicalClusterRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 待重启的逻辑集群的ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterId string `json:"logical_cluster_id"`
}

func (o RestartLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestartLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"RestartLogicalClusterRequest", string(data)}, " ")
}
