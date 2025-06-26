package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceReplicationPoliciesResponse Response Object
type ListInstanceReplicationPoliciesResponse struct {

	// 同步策略列表
	Policies *[]ReplicationPolicy `json:"policies,omitempty"`

	// 同步策略总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceReplicationPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceReplicationPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceReplicationPoliciesResponse", string(data)}, " ")
}
