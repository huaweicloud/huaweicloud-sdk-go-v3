package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TrendRiskInfo struct {

	// **参数解释**: 基线风险数量 **取值范围**: 最小值0，最大值2147483647
	BaselineRiskNum *int32 `json:"baseline_risk_num,omitempty"`

	// **参数解释**: 资产风险数量 **取值范围**: 最小值0，最大值2147483647
	AssetRiskNum *int32 `json:"asset_risk_num,omitempty"`

	// **参数解释**: 告警数量 **取值范围**: 最小值0，最大值2147483647
	AlarmRiskNum *int32 `json:"alarm_risk_num,omitempty"`

	// **参数解释**: 漏洞数 **取值范围**: 最小值0，最大值2147483647
	VulRiskNum *int32 `json:"vul_risk_num,omitempty"`

	// **参数解释**： 创建时间 **取值范围**： 最小值0，最大值9223372036854775807
	UpdateTime *int64 `json:"update_time,omitempty"`

	// **参数解释**: 时间，毫秒级时间戳 **取值范围**: 最小值0，最大值9223372036854775807
	Time *int64 `json:"time,omitempty"`

	// **参数解释**: 分数 **取值范围**: 最小值0，最大值100
	Score *float32 `json:"score,omitempty"`
}

func (o TrendRiskInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrendRiskInfo struct{}"
	}

	return strings.Join([]string{"TrendRiskInfo", string(data)}, " ")
}
