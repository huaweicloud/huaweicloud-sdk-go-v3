package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueRequest Request Object
type ListWorkloadQueueRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	// 逻辑集群名称。逻辑集群模式下该字段必填。
	LogicalClusterName *string `json:"logical_cluster_name,omitempty"`
}

func (o ListWorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueRequest", string(data)}, " ")
}
