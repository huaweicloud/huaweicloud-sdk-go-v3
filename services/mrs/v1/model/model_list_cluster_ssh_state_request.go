package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterSshStateRequest Request Object
type ListClusterSshStateRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterSshStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterSshStateRequest struct{}"
	}

	return strings.Join([]string{"ListClusterSshStateRequest", string(data)}, " ")
}
