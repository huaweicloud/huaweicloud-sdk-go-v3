package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PolicyAdvancedRetentionRules 按照时间的高级保留策略
type PolicyAdvancedRetentionRules struct {
	WeeklyRetentionRules *PolicyWeeklyRetentionRules `json:"weekly_retention_rules,omitempty"`

	MonthlyRetentionRules *PolicyMonthlyRetentionRules `json:"monthly_retention_rules,omitempty"`

	YearlyRetentionRules *PolicyYearlyRetentionRules `json:"yearly_retention_rules,omitempty"`
}

func (o PolicyAdvancedRetentionRules) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PolicyAdvancedRetentionRules struct{}"
	}

	return strings.Join([]string{"PolicyAdvancedRetentionRules", string(data)}, " ")
}
