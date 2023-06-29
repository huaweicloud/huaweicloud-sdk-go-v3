package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AddWorkloadQueueRequest Request Object
type AddWorkloadQueueRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`

	Body *WorkloadQueueReq `json:"body,omitempty"`
}

func (o AddWorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AddWorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"AddWorkloadQueueRequest", string(data)}, " ")
}
