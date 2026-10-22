package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DisassociateAutoSnapshotPolicyRequest Request Object
type DisassociateAutoSnapshotPolicyRequest struct {

	// 自动快照策略ID。
	SnapshotPolicyId string `json:"snapshot_policy_id"`

	Body *[]string `json:"body,omitempty"`
}

func (o DisassociateAutoSnapshotPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DisassociateAutoSnapshotPolicyRequest struct{}"
	}

	return strings.Join([]string{"DisassociateAutoSnapshotPolicyRequest", string(data)}, " ")
}
