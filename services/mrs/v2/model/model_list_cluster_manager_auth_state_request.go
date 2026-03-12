package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterManagerAuthStateRequest Request Object
type ListClusterManagerAuthStateRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterManagerAuthStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterManagerAuthStateRequest struct{}"
	}

	return strings.Join([]string{"ListClusterManagerAuthStateRequest", string(data)}, " ")
}
