package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterDetailsRequest Request Object
type ListClusterDetailsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterDetailsRequest", string(data)}, " ")
}
