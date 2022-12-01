package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateClusterCnRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	Body *BatchCreateCn `json:"body,omitempty"`
}

func (o BatchCreateClusterCnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateClusterCnRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateClusterCnRequest", string(data)}, " ")
}
