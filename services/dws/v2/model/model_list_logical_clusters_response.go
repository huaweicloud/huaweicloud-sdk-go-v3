package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListLogicalClustersResponse Response Object
type ListLogicalClustersResponse struct {

	// 逻辑集群列表信息
	LogicalClusters *[]LogicalClusterInfo `json:"logical_clusters,omitempty"`

	// 逻辑集群总数量
	Count *int32 `json:"count,omitempty"`

	// 作为互斥结果，如果集群内有其他运维操作，该值为false，此时不能添加逻辑集群
	AddEnable      *bool `json:"add_enable,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ListLogicalClustersResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLogicalClustersResponse struct{}"
	}

	return strings.Join([]string{"ListLogicalClustersResponse", string(data)}, " ")
}
