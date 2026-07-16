package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowClusterRequest Request Object
type ShowClusterRequest struct {

	// **参数解释**：纳管集群ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ClusterId string `json:"cluster_id"`
}

func (o ShowClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowClusterRequest struct{}"
	}

	return strings.Join([]string{"ShowClusterRequest", string(data)}, " ")
}
