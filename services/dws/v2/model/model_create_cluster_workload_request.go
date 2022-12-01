package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateClusterWorkloadRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *WorkloadStatus `json:"body,omitempty"`
}

func (o CreateClusterWorkloadRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClusterWorkloadRequest struct{}"
	}

	return strings.Join([]string{"CreateClusterWorkloadRequest", string(data)}, " ")
}
