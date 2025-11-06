package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupStrategyDetail **参数解释**： 备份策略详细信息。 **取值范围**： 不涉及。
type BackupStrategyDetail struct {

	// **参数解释**： 策略ID。 **取值范围**： 不涉及。
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**： 策略名称。 **取值范围**： 不涉及。
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 执行策略，一般为crontab表达式。 **取值范围**： 不涉及。
	BackupStrategy *string `json:"backup_strategy,omitempty"`

	// **参数解释**： 备份类型。 **取值范围**： 不涉及。
	BackupType *string `json:"backup_type,omitempty"`

	// **参数解释**： 备份级别。 **取值范围**： 不涉及。
	BackupLevel *string `json:"backup_level,omitempty"`

	// **参数解释**： 下次触发时间（预估，与其它任务冲突时不执行）。 **取值范围**： 不涉及。
	NextFireTime *string `json:"next_fire_time,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 时区偏移量（相比UTC时间）。 **取值范围**： 0~23
	TimeZoneOffset *int32 `json:"time_zone_offset,omitempty"`

	// **参数解释**： 备份的数据库。 **取值范围**： 不涉及。
	BackupDatabase *string `json:"backup_database,omitempty"`

	// **参数解释**： 备份的数据库模式列表。 **取值范围**： 不涉及。
	BackupSchemaList *string `json:"backup_schema_list,omitempty"`

	// **参数解释**： 备份的数据库表列表。 **取值范围**： 不涉及。
	BackupTableList *string `json:"backup_table_list,omitempty"`
}

func (o BackupStrategyDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyDetail struct{}"
	}

	return strings.Join([]string{"BackupStrategyDetail", string(data)}, " ")
}
