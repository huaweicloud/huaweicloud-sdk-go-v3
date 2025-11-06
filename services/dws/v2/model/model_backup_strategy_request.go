package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BackupStrategyRequest **参数解释**： 备份策略请求信息。 **取值范围**： 不涉及。
type BackupStrategyRequest struct {

	// **参数解释**： 策略ID。 **取值范围**： 不涉及。
	PolicyId *string `json:"policy_id,omitempty"`

	// **参数解释**： 策略名称。添加备份策略时为必选字段。策略名称在4位到92位之间，必须以字母开头，不区分大小写，可以包含字母、数字、中划线或者下划线，不能包含其他特殊字符，并且名称唯一。 **取值范围**： 不涉及。
	PolicyName *string `json:"policy_name,omitempty"`

	// **参数解释**： 执行策略。添加备份策略时为必选字段。符合Cron表达式格式。 **取值范围**： 不涉及。
	BackupStrategy *string `json:"backup_strategy,omitempty"`

	// **参数解释**： 备份类型。 **取值范围**： full：全量。 increment：增量。
	BackupType *string `json:"backup_type,omitempty"`

	// **参数解释**： 备份级别。 **取值范围**： cluster：集群级。 schema：模式级。 table：表级。
	BackupLevel *string `json:"backup_level,omitempty"`

	// **参数解释**： 下一次触发时间。 **取值范围**： 不涉及。
	NextFireTime *string `json:"next_fire_time,omitempty"`

	// **参数解释**： 更新时间。 **取值范围**： 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 时区偏移量。 **取值范围**： 不涉及。
	TimeZoneOffset *int32 `json:"time_zone_offset,omitempty"`

	// **参数解释**： 备份的数据库。 **取值范围**： 不涉及。
	BackupDatabase *string `json:"backup_database,omitempty"`

	// **参数解释**： 备份的数据库模式列表。 **取值范围**： 不涉及。
	BackupSchemaList *string `json:"backup_schema_list,omitempty"`

	// **参数解释**： 备份的数据库表列表。 **取值范围**： 不涉及。
	BackupTableList *string `json:"backup_table_list,omitempty"`
}

func (o BackupStrategyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupStrategyRequest struct{}"
	}

	return strings.Join([]string{"BackupStrategyRequest", string(data)}, " ")
}
