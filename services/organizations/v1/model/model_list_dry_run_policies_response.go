package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDryRunPoliciesResponse Response Object
type ListDryRunPoliciesResponse struct {

	// 组织中的策略列表。
	Policies *[]PolicySummaryDto `json:"policies,omitempty"`

	PageInfo       *PageInfoDto `json:"page_info,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListDryRunPoliciesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDryRunPoliciesResponse struct{}"
	}

	return strings.Join([]string{"ListDryRunPoliciesResponse", string(data)}, " ")
}
