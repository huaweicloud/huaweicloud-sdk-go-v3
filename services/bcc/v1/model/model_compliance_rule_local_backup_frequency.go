package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ComplianceRuleLocalBackupFrequency 本地备份频率，如按天 2次/天，按周 周一，三，五每天执行2次
type ComplianceRuleLocalBackupFrequency struct {

	// 备份频率类型，可选daily weekly
	Type *string `json:"type,omitempty"`

	// 每日执行次数
	TimesPerDay int32 `json:"times_per_day"`

	// daily类型下，间隔多少天，默认值：1
	Interval *int32 `json:"interval,omitempty"`

	// 如果type是weekly必选，表示一周内那些天备份 如：[\"MO\", \"TU\", \"WE\", \"TH\", \"FR\", \"SA\", \"SU\"]
	DaysOfWeek *[]string `json:"days_of_week,omitempty"`
}

func (o ComplianceRuleLocalBackupFrequency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ComplianceRuleLocalBackupFrequency struct{}"
	}

	return strings.Join([]string{"ComplianceRuleLocalBackupFrequency", string(data)}, " ")
}
