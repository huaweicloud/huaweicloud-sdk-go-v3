package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AssociateAutoSnapshotPolicyRequest Request Object
type AssociateAutoSnapshotPolicyRequest struct {

	// 自动快照策略ID。
	SnapshotPolicyId string `json:"snapshot_policy_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o AssociateAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AssociateAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"AssociateAutoSnapshotPolicyRequest", string(data)}, " ")
}
