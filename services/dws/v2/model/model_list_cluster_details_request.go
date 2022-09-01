package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClusterDetailsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id" xml:"cluster_id"`
}

func (o ListClusterDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterDetailsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterDetailsRequest", string(data)}, " ")
}
