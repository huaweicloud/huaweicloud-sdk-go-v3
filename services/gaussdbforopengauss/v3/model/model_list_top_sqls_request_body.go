package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTopSqlsRequestBody struct {

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 所选实例节点ID列表。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeIds []string `json:"node_ids"`

	// **参数解释**: 起始时间。 **约束限制**: 13位UNIX时间戳格式，单位是毫秒，时区是UTC。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**: 结束时间。 **约束限制**: 13位UNIX时间戳格式，单位是毫秒，时区是UTC。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**: 起始时间。 **约束限制**: UTC时间。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTimeUtc *string `json:"start_time_utc,omitempty"`

	// **参数解释**: 结束时间。 **约束限制**: UTC时间。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTimeUtc *string `json:"end_time_utc,omitempty"`

	// **参数解释**: Top SQL的用户名。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**: Top SQL的SQL文本，例如 select pg_sleep(5)。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlText *string `json:"sql_text,omitempty"`

	// **参数解释**: Top SQL的归一化SQL ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: Top SQL的数据库名。 **约束限制**: 引擎版本8.200及以上显示。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	DbName *string `json:"db_name,omitempty"`

	// **参数解释**: 是否支持展示系统用户。 **约束限制**: 不涉及。 **取值范围**: - true：支持展示系统用户。 - false：不支持展示系统用户。  **默认取值**: 不涉及
	SupportSystem *bool `json:"support_system,omitempty"`

	// **参数解释**: 字段汇聚查询条件列表。 **约束限制**: 只支持针对query字段全与或者全或的查询。
	MultiQueries *[]MultiQueryConditionOption `json:"multi_queries,omitempty"`
}

func (o ListTopSqlsRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopSqlsRequestBody struct{}"
	}

	return strings.Join([]string{"ListTopSqlsRequestBody", string(data)}, " ")
}
