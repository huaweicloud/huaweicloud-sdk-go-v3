package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityCheckStatisticResponse Response Object
type ShowSecurityCheckStatisticResponse struct {

	// **参数解释**: 本次体检服务器数量 **取值范围**: 最小值0，最大值2147483647
	TotalHostNum *int32 `json:"total_host_num,omitempty"`

	// **参数解释**: 有安全风险的服务器数量 **取值范围**: 最小值0，最大值2147483647
	RiskHostNum *int32 `json:"risk_host_num,omitempty"`

	// **参数解释**: 高危风险的服务器数量 **取值范围**: 最小值0，最大值2147483647
	HighRiskHostNum *int32 `json:"high_risk_host_num,omitempty"`

	// **参数解释**: 中危风险的服务器数量 **取值范围**: 最小值0，最大值2147483647
	MediumRiskHostNum *int32 `json:"medium_risk_host_num,omitempty"`

	// **参数解释**: 低危风险的服务器数量 **取值范围**: 最小值0，最大值2147483647
	LowRiskHostNum *int32 `json:"low_risk_host_num,omitempty"`

	// **参数解释**: 无风险的服务器数量 **取值范围**: 最小值0，最大值2147483647
	SafeHostNum    *int32 `json:"safe_host_num,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSecurityCheckStatisticResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityCheckStatisticResponse struct{}"
	}

	return strings.Join([]string{"ShowSecurityCheckStatisticResponse", string(data)}, " ")
}
