package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregatePolicyStateComplianceSummaryRequest Request Object
type ShowAggregatePolicyStateComplianceSummaryRequest struct {
	Body *AggregatePolicyStatesRequest `json:"body,omitempty"`
}

func (o ShowAggregatePolicyStateComplianceSummaryRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregatePolicyStateComplianceSummaryRequest struct{}"
	}

	return strings.Join([]string{"ShowAggregatePolicyStateComplianceSummaryRequest", string(data)}, " ")
}
