package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListClusterSnapshotsRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListClusterSnapshotsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterSnapshotsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterSnapshotsRequest", string(data)}, " ")
}
