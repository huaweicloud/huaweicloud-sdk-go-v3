package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SlowSqlDetailRequestBody struct {

	// **参数解释**: 起始日期。 **约束限制**: 格式:13位UNIX时间戳格式，单位是毫秒，时区是UTC。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	StartTime int64 `json:"start_time"`

	// **参数解释**: 结束日期。 **约束限制**: 格式:13位UNIX时间戳格式，单位是毫秒，时区是UTC。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	EndTime int64 `json:"end_time"`

	// **参数解释**: 实例ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**: 字符型组合过滤多条件查询列表。 **约束限制**: 不涉及。
	MultiQueries *[]MultiQueryConditionOption `json:"multi_queries,omitempty"`

	// **参数解释**: 慢SQL的SQL ID。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	SqlId string `json:"sql_id"`

	// **参数解释**: 节点ID数组。 **约束限制**: 不涉及。
	NodeIds []string `json:"node_ids"`
}

func (o SlowSqlDetailRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SlowSqlDetailRequestBody struct{}"
	}

	return strings.Join([]string{"SlowSqlDetailRequestBody", string(data)}, " ")
}
