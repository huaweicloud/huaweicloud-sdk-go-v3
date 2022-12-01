package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteSnapshotPolicyRequest struct {

	// 集群ID。
	ClusterId string `json:"cluster_id"`

	// 快照策略ID。
	Id string `json:"id"`
}

func (o DeleteSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSnapshotPolicyRequest", string(data)}, " ")
}
