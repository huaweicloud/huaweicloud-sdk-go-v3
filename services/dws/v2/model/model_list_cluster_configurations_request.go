package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClusterConfigurationsRequest struct {

	// 集群的ID。
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterConfigurationsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterConfigurationsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterConfigurationsRequest", string(data)}, " ")
}
