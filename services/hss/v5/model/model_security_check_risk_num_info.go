package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SecurityCheckRiskNumInfo 风险数量信息
type SecurityCheckRiskNumInfo struct {

	// **参数解释**： 风险总数 **取值范围**： 不涉及
	TotalRiskNum *int64 `json:"total_risk_num,omitempty"`

	// **参数解释**： 与上一周期比较的状态 **取值范围**： - lower：降低 - increase：增加 - equals：持平
	CompareStatus *string `json:"compare_status,omitempty"`

	// **参数解释**： 与上一周期相差的数量的绝对值 **取值范围**： 不涉及
	CompareNum *int32 `json:"compare_num,omitempty"`

	// **参数解释**： 高危风险数量 **取值范围**： 不涉及
	HighRiskNum *int32 `json:"high_risk_num,omitempty"`
}

func (o SecurityCheckRiskNumInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SecurityCheckRiskNumInfo struct{}"
	}

	return strings.Join([]string{"SecurityCheckRiskNumInfo", string(data)}, " ")
}
