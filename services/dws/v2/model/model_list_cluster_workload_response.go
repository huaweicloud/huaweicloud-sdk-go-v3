package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListClusterWorkloadResponse struct {
	WorkloadStatus *WorkloadStatus `json:"workload_status,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListClusterWorkloadResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterWorkloadResponse struct{}"
	}

	return strings.Join([]string{"ListClusterWorkloadResponse", string(data)}, " ")
}
