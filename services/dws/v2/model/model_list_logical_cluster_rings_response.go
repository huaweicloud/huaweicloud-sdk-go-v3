package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClusterRingsResponse Response Object
type ListLogicalClusterRingsResponse struct {

	// 集群环列表信息
	ClusterRings *[]LogicalClusterRingInfo `json:"cluster_rings,omitempty"`

	// 集群环数量
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
