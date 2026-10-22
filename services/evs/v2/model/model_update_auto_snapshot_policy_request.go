package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateAutoSnapshotPolicyRequest Request Object
type UpdateAutoSnapshotPolicyRequest struct {

	// 自动快照策略ID。
	SnapshotPolicyId string `json:"snapshot_policy_id"`

	Body *UpdateSnapshotPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateAutoSnapshotPolicyRequest", string(data)}, " ")
}
