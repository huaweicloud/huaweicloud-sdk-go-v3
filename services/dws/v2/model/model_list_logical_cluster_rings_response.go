package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterRingsResponse Response Object
type ListLogicalClusterRingsResponse struct {

	// **参数解释**： 集群环列表信息。 **取值范围**： 不涉及。
	ClusterRings *[]LogicalClusterRingInfo `json:"cluster_rings,omitempty"`

	// **参数解释**： 集群环数量。 **取值范围**： 不涉及。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListLogicalClusterRingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClusterRingsResponse struct{}"
	}

	return strings.Join([]string{"ListLogicalClusterRingsResponse", string(data)}, " ")
}
