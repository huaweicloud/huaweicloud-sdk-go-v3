package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoSnapshotPolicyResponse Response Object
type ListAutoSnapshotPolicyResponse struct {

	// 查询的自动快照策略列表信息。
	SnapshotPolicies *[]GetSnapshotPolicy `json:"snapshot_policies,omitempty"`
	HttpStatusCode   int                  `json:"-"`
}

func (o ListAutoSnapshotPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoSnapshotPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListAutoSnapshotPolicyResponse", string(data)}, " ")
}
