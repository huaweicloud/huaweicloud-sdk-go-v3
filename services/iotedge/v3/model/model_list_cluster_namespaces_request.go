package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterNamespacesRequest Request Object
type ListClusterNamespacesRequest struct {

	// 边缘集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterNamespacesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterNamespacesRequest struct{}"
	}

	return strings.Join([]string{"ListClusterNamespacesRequest", string(data)}, " ")
}
