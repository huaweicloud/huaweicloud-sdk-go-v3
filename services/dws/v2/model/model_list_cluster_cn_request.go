package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClusterCnRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterCnRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterCnRequest struct{}"
	}

	return strings.Join([]string{"ListClusterCnRequest", string(data)}, " ")
}
