package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OperationDefinitionRequestInfo **参数解释**: 调度参数 **约束限制**: 不涉及 **取值范围**: 取值为0到100。 **默认取值**: 不涉及
type OperationDefinitionRequestInfo struct {

	// **参数解释**: 保留日备个数，该备份不受保留最大备份数限制。若选择该参数，则timezone也必选。 **约束限制**: 不涉及 **取值范围**: 取值为0到100。 **默认取值**: 不涉及
	DayBackups *int32 `json:"day_backups,omitempty"`

	// **参数解释**: 单个备份对象自动备份的最大备份数。-1代表不按备份数清理。若该字段和retention_duration_days字段同时为空，备份会永久保留。 **约束限制**: 不涉及 **取值范围**: 取值为-1或1-99999 **默认取值**: -1
	MaxBackups *int32 `json:"max_backups,omitempty"`

	// **参数解释**: 保留月备个数，该备份不受保留最大备份数限制。若选择该参数，则timezone也必选。 **约束限制**: 不涉及 **取值范围**: 取值为0到100。 **默认取值**: 不涉及
	MonthBackups *int32 `json:"month_backups,omitempty"`

	// **参数解释**: 备份保留时长，单位天。最长支持99999天。-1代表不按时间清理。若该字段和max_backups参数同时为空，备份会永久保留。 **约束限制**: 不涉及 **取值范围**: 取值范围-1-99999 **默认取值**: -1
	RetentionDurationDays *int32 `json:"retention_duration_days,omitempty"`

	// **参数解释**: 用户所在时区，格式形如UTC+08:00 **约束限制**: 若没有选择年备，月备，周备，日备中任一参数，则不能选择该参数。 **取值范围**: - UTC+05:00 - UTC+06:00 - UTC+07:00 ... **默认取值**: 不涉及
	Timezone *string `json:"timezone,omitempty"`

	// **参数解释**: 保留周备个数，该备份不受保留最大备份数限制。若选择该参数，则timezone也必选。 **约束限制**: 不涉及 **取值范围**: 取值为0到100。 **默认取值**: 不涉及
	WeekBackups *int32 `json:"week_backups,omitempty"`

	// **参数解释**: 保留年备个数，该备份不受保留最大备份数限制。若选择该参数，则timezone也必选。 **约束限制**: 不涉及 **取值范围**: 取值为0到100。 **默认取值**: 不涉及
	YearBackups *int32 `json:"year_backups,omitempty"`
}

func (o OperationDefinitionRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OperationDefinitionRequestInfo struct{}"
	}

	return strings.Join([]string{"OperationDefinitionRequestInfo", string(data)}, " ")
}
