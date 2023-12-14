package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopoRingsResponse Response Object
type ListTopoRingsResponse struct {

	// 集群拓扑环列表信息
	ClusterRings *[]TopoRingInfo `json:"cluster_rings,omitempty"`

	// 集群环数量
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTopoRingsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopoRingsResponse struct{}"
	}

	return strings.Join([]string{"ListTopoRingsResponse", string(data)}, " ")
}
