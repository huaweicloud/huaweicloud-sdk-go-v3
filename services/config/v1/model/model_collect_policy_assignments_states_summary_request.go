package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CollectPolicyAssignmentsStatesSummaryRequest Request Object
type CollectPolicyAssignmentsStatesSummaryRequest struct {

	// 规则ID
	PolicyAssignmentId string `json:"policy_assignment_id"`

	// 资源名称
	ResourceName *string `json:"resource_name,omitempty"`
}

func (o CollectPolicyAssignmentsStatesSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectPolicyAssignmentsStatesSummaryRequest struct{}"
	}

	return strings.Join([]string{"CollectPolicyAssignmentsStatesSummaryRequest", string(data)}, " ")
}
