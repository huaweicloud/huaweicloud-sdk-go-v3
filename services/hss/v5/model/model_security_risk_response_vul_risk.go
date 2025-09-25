package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRiskResponseVulRisk **参数解释**： 漏洞风险
type SecurityRiskResponseVulRisk struct {

	// **参数解释**： 漏洞风险列表 **取值范围**: 最小值0，最大值4
	RiskList *[]RiskListInfo `json:"risk_list,omitempty"`

	// **参数解释**： 扣分 **取值范围**: 最小值0，最大值20
	DeductScore *float32 `json:"deduct_score,omitempty"`

	// **参数解释**： 未进行漏洞扫描主机数(一个月内) **取值范围**: 最小值0，最大值2147483647
	UnScannedHostNum *int32 `json:"un_scanned_host_num,omitempty"`

	// **参数解释**： 风险总数 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum *int32 `json:"total_risk_num,omitempty"`
}

func (o SecurityRiskResponseVulRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRiskResponseVulRisk struct{}"
	}

	return strings.Join([]string{"SecurityRiskResponseVulRisk", string(data)}, " ")
}
