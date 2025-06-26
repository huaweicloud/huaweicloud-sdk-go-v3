package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteLogicalClusterRequest Request Object
type DeleteLogicalClusterRequest struct {

	// **参数解释**： 集群ID。获取方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 指定待删除逻辑集群的ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterId string `json:"logical_cluster_id"`
}

func (o DeleteLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteLogicalClusterRequest", string(data)}, " ")
}
