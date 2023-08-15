package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPolicyAssignmentsRequest Request Object
type ListPolicyAssignmentsRequest struct {
}

func (o ListPolicyAssignmentsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPolicyAssignmentsRequest struct{}"
	}

	return strings.Join([]string{"ListPolicyAssignmentsRequest", string(data)}, " ")
}
