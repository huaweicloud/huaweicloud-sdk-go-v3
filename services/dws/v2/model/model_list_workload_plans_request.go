package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadPlansRequest Request Object
type ListWorkloadPlansRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 逻辑集群名称
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`

	// 查询条数
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListWorkloadPlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadPlansRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadPlansRequest", string(data)}, " ")
}
