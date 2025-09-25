package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListSqlLimitTaskRequestBody struct {

	// **参数解释**: 限流任务范围。 **约束限制**: 不涉及。 **取值范围**: 目前支持SQL、SESSION。 **默认取值**: 不涉及。
	TaskScope *string `json:"task_scope,omitempty"`

	// **参数解释**: 限流类型。 **约束限制**: 不涉及。 **取值范围**: 支持SQL_ID、SQL_TYPE、SESSION_ACTIVE_MAX_COUNT类型。 **默认取值**: 不涉及。
	LimitType *string `json:"limit_type,omitempty"`

	// **参数解释**: 限流类型值，支持模糊匹配。 **约束限制**: 不涉及。 **取值范围**: 长度为1~19字符，且只能包含大小写字母、数字和_。 **默认取值**: 不涉及。
	LimitTypeValue *string `json:"limit_type_value,omitempty"`

	// **参数解释**: 限流任务名，支持模糊匹配。 **约束限制**: 不涉及。 **取值范围**: 长度为1~100字符，只能包含大小写字母、数字、-、_和$。 **默认取值**: 不涉及。
	TaskName *string `json:"task_name,omitempty"`

	// **参数解释**: 规则名。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	RuleName *string `json:"rule_name,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 不涉及。 **取值范围**: 0 - 10000 **默认取值**: 0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 查询记录数。 **约束限制**: 不涉及。 **取值范围**: 不能为负数，最小值为1，最大值为100。 **默认取值**: 10
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 限流任务SQL_ID。 **约束限制**: 不涉及。 **取值范围**: 1 到 19 位的正整数（首位不为 0）。 **默认取值**: 不涉及。
	SqlId *string `json:"sql_id,omitempty"`

	// **参数解释**: 限流任务NodeID列表。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeIds *[]string `json:"node_ids,omitempty"`
}

func (o ListSqlLimitTaskRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlLimitTaskRequestBody struct{}"
	}

	return strings.Join([]string{"ListSqlLimitTaskRequestBody", string(data)}, " ")
}
