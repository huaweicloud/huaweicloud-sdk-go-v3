package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteAutoSnapshotPolicyRequest Request Object
type DeleteAutoSnapshotPolicyRequest struct {

	// 自动快照策略ID。
	SnapshotPolicyId string `json:"snapshot_policy_id"`
}

func (o DeleteAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteAutoSnapshotPolicyRequest", string(data)}, " ")
}
