package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRiskResponseBaselineRisk **参数解释**： 基线风险
type SecurityRiskResponseBaselineRisk struct {

	// **参数解释**： 基线风险列表 **取值范围**: 最小值0，最大值4
	RiskList *[]RiskListInfo `json:"risk_list,omitempty"`

	// **参数解释**： 扣分 **取值范围**: 最小值0，最大值20
	DeductScore *float32 `json:"deduct_score,omitempty"`

	// **参数解释**： 未开启的策略列表 **取值范围**: 最小值0，最大值32
	PolicyList *[]PolicyInfo `json:"policy_list,omitempty"`

	// **参数解释**： 存在弱口令主机数量 **取值范围**: 最小值0，最大值2147483647
	ExistedPwdHostNum *int32 `json:"existed_pwd_host_num,omitempty"`

	// **参数解释**： 未进行基线配置检查的主机数 **取值范围**: 最小值0，最大值2147483647
	UnScannedBaselineHostNum *int32 `json:"un_scanned_baseline_host_num,omitempty"`

	// **参数解释**： 风险总数 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum *int32 `json:"total_risk_num,omitempty"`
}

func (o SecurityRiskResponseBaselineRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRiskResponseBaselineRisk struct{}"
	}

	return strings.Join([]string{"SecurityRiskResponseBaselineRisk", string(data)}, " ")
}
