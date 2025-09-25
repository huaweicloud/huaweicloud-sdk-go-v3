package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityRiskResponseAssetRisk **参数解释**： 资产风险
type SecurityRiskResponseAssetRisk struct {

	// **参数解释**： 存在危险端口主机数 **取值范围**: 最小值0，最大值2147483647
	ExistedDangerPortHostNum *int32 `json:"existed_danger_port_host_num,omitempty"`

	// **参数解释**： 策略列表 **取值范围**: 最小值0，最大值32
	PolicyList *[]PolicyInfo `json:"policy_list,omitempty"`

	// **参数解释**： 扣分 **取值范围**: 最小值0，最大值10
	DeductScore *float32 `json:"deduct_score,omitempty"`

	// **参数解释**： 风险总数 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum *int32 `json:"total_risk_num,omitempty"`
}

func (o SecurityRiskResponseAssetRisk) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityRiskResponseAssetRisk struct{}"
	}

	return strings.Join([]string{"SecurityRiskResponseAssetRisk", string(data)}, " ")
}
