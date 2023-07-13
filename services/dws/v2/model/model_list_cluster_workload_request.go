package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterWorkloadRequest Request Object
type ListClusterWorkloadRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterWorkloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterWorkloadRequest struct{}"
	}

	return strings.Join([]string{"ListClusterWorkloadRequest", string(data)}, " ")
}
