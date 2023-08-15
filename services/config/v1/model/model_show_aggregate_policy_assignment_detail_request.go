package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregatePolicyAssignmentDetailRequest Request Object
type ShowAggregatePolicyAssignmentDetailRequest struct {
	Body *AggregatePolicyAssignmentDetailRequest `json:"body,omitempty"`
}

func (o ShowAggregatePolicyAssignmentDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregatePolicyAssignmentDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowAggregatePolicyAssignmentDetailRequest", string(data)}, " ")
}
