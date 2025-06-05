package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteComputingClusterRequest Request Object
type DeleteComputingClusterRequest struct {

	// 计算集群ID。
	ClusterId string `json:"cluster_id"`
}

func (o DeleteComputingClusterRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteComputingClusterRequest struct{}"
	}

	return strings.Join([]string{"DeleteComputingClusterRequest", string(data)}, " ")
}
