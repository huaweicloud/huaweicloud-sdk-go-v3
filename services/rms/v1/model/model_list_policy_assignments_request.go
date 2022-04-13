package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListPolicyAssignmentsRequest struct {
}

func (o ListPolicyAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyAssignmentsRequest", string(data)}, " ")
}
