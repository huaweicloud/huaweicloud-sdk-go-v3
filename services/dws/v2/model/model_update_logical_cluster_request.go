package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterRequest Request Object
type UpdateLogicalClusterRequest struct {

	// **参数解释**： 集群ID。获取方式方法请参见[获取集群ID](dws_02_00068.xml)。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ClusterId string `json:"cluster_id"`

	// **参数解释**： 指定待编辑逻辑集群的ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LogicalClusterId string `json:"logical_cluster_id"`

	Body *UpdateLogicalClusterRequestBody `json:"body,omitempty"`
}

func (o UpdateLogicalClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterRequest struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterRequest", string(data)}, " ")
}
