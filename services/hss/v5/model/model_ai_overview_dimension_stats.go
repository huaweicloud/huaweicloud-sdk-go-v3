package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AiOverviewDimensionStats AI Defender防护维度统计项
type AiOverviewDimensionStats struct {

	// **参数解释**: 防护维度key，取值及含义如下： - consistency_detection：意图行为一致性检测 - command_control：命令执行控制 - file_access_control：文件访问控制 - sensitive_info_detection：敏感信息检测 - role_limit：角色限定（行为轨迹检测不覆盖此维度） - openclaw_detectors：记忆投毒防护（OpenClaw专属维度） - other：其他类  **取值范围**: consistency_detection、command_control、file_access_control、sensitive_info_detection、role_limit、openclaw_detectors、other
	Dimension *string `json:"dimension,omitempty"`

	// **参数解释**： 告警数量，type=alarm时返回 **取值范围**： 最小值0，最大值9223372036854775807
	AlarmCount *int64 `json:"alarm_count,omitempty"`

	// **参数解释**： 风险数量，type=risk时返回，role_limit固定为0 **取值范围**： 最小值0，最大值9223372036854775807
	RiskCount *int64 `json:"risk_count,omitempty"`
}

func (o AiOverviewDimensionStats) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AiOverviewDimensionStats struct{}"
	}

	return strings.Join([]string{"AiOverviewDimensionStats", string(data)}, " ")
}
