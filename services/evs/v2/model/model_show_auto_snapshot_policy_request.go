package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAutoSnapshotPolicyRequest Request Object
type ShowAutoSnapshotPolicyRequest struct {

	// 自动快照策略ID。
	SnapshotPolicyId string `json:"snapshot_policy_id"`
}

func (o ShowAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowAutoSnapshotPolicyRequest", string(data)}, " ")
}
