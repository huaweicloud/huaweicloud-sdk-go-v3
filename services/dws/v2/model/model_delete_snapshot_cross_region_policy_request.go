package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSnapshotCrossRegionPolicyRequest Request Object
type DeleteSnapshotCrossRegionPolicyRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o DeleteSnapshotCrossRegionPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotCrossRegionPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotCrossRegionPolicyRequest", string(data)}, " ")
}
