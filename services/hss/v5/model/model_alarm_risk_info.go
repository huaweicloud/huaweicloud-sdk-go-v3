package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmRiskInfo struct {

	// **参数解释**： 风险列表 **取值范围**: 最小值0，最大值4
	RiskList *[]RiskListInfo `json:"risk_list,omitempty"`

	// **参数解释**： 扣分 **取值范围**: 最小值0，最大值30
	DeductScore *float32 `json:"deduct_score,omitempty"`

	// **参数解释**：  策略信息  **取值范围**:  最小值0，最大值32
	PolicyList *[]PolicyInfo `json:"policy_list,omitempty"`

	// **参数解释**： 风险总数 **取值范围**: 最小值0，最大值2147483647
	TotalRiskNum *int32 `json:"total_risk_num,omitempty"`
}

func (o AlarmRiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRiskInfo struct{}"
	}

	return strings.Join([]string{"AlarmRiskInfo", string(data)}, " ")
}
