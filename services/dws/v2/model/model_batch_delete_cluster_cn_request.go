package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteClusterCnRequest Request Object
type BatchDeleteClusterCnRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`

	Body *BatchDeleteCn `json:"body,omitempty"`
}

func (o BatchDeleteClusterCnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteClusterCnRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteClusterCnRequest", string(data)}, " ")
}
