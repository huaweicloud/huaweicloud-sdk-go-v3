package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAggregatePolicyStateComplianceSummaryResponse Response Object
type ShowAggregatePolicyStateComplianceSummaryResponse struct {

	// 合规总结结果列表
	Results        *[]AggregatePolicyComplianceSummaryResult `json:"results,omitempty"`
	HttpStatusCode int                                       `json:"-"`
}

func (o ShowAggregatePolicyStateComplianceSummaryResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAggregatePolicyStateComplianceSummaryResponse struct{}"
	}

	return strings.Join([]string{"ShowAggregatePolicyStateComplianceSummaryResponse", string(data)}, " ")
}
