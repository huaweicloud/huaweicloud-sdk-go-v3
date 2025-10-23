package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateComplianceRuleRequestBodyLocalBackupFrequency 本地备份频率，如按天 2次/天，按周 周一，三，五每天执行2次
type UpdateComplianceRuleRequestBodyLocalBackupFrequency struct {

	// 备份频率类型，可选daily weekly
	Type *string `json:"type,omitempty"`

	// daily类型下，间隔多少天，默认值：1
	Interval *int32 `json:"interval,omitempty"`

	// 每日执行次数
	TimesPerDay int32 `json:"times_per_day"`

	// 如果type是weekly必选，表示一周内那些天备份 如：[\"MO\", \"TU\", \"WE\", \"TH\", \"FR\", \"SA\", \"SU\"]
	DaysOfWeek *[]string `json:"days_of_week,omitempty"`
}

func (o UpdateComplianceRuleRequestBodyLocalBackupFrequency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateComplianceRuleRequestBodyLocalBackupFrequency struct{}"
	}

	return strings.Join([]string{"UpdateComplianceRuleRequestBodyLocalBackupFrequency", string(data)}, " ")
}
