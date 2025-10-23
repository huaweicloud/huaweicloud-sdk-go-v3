package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QuerySqlPlanStateRequest struct {

	// **参数解释**: 归一化的SQL ID **约束限制**: 不涉及。
	SqlId string `json:"sql_id"`

	// **参数解释**: SQL执行计划每页显示数量。 **约束限制**: 不涉及。 **取值范围**: 整数，1~100。 **默认取值**: 不涉及。
	PageSize string `json:"page_size"`

	// **参数解释**: SQL执行计划起始页码。 **约束限制**: 不涉及。 **取值范围**: 大于等于0的整数。 **默认取值**: 不涉及。
	Offset string `json:"offset"`
}

func (o QuerySqlPlanStateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QuerySqlPlanStateRequest struct{}"
	}

	return strings.Join([]string{"QuerySqlPlanStateRequest", string(data)}, " ")
}
