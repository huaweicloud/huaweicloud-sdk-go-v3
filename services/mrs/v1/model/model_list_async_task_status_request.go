package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAsyncTaskStatusRequest Request Object
type ListAsyncTaskStatusRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListAsyncTaskStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAsyncTaskStatusRequest struct{}"
	}

	return strings.Join([]string{"ListAsyncTaskStatusRequest", string(data)}, " ")
}
