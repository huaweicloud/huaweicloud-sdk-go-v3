package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountAutoSnapshotPoliciesResponse Response Object
type CountAutoSnapshotPoliciesResponse struct {

	// 自动快照策略的数量。
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountAutoSnapshotPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountAutoSnapshotPoliciesResponse struct{}"
	}

	return strings.Join([]string{"CountAutoSnapshotPoliciesResponse", string(data)}, " ")
}
