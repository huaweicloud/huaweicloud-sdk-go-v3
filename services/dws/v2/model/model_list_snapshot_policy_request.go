package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSnapshotPolicyRequest Request Object
type ListSnapshotPolicyRequest struct {

	// 集群ID
	ClusterId string `json:"cluster_id"`
}

func (o ListSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"ListSnapshotPolicyRequest", string(data)}, " ")
}
