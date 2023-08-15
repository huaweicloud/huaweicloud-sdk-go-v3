package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkloadQueueRequest Request Object
type ListWorkloadQueueRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListWorkloadQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkloadQueueRequest struct{}"
	}

	return strings.Join([]string{"ListWorkloadQueueRequest", string(data)}, " ")
}
