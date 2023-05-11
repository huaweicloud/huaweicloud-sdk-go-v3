package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowAggregateComplianceDetailsByPolicyAssignmentResponse struct {

	// 合规结果查询返回值
	PolicyStates *[]PolicyState `json:"policy_states,omitempty"`

	PageInfo       *PageInfo `json:"page_info,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ShowAggregateComplianceDetailsByPolicyAssignmentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregateComplianceDetailsByPolicyAssignmentResponse struct{}"
	}

	return strings.Join([]string{"ShowAggregateComplianceDetailsByPolicyAssignmentResponse", string(data)}, " ")
}
