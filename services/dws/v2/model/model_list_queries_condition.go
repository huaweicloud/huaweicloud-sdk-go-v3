package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListQueriesCondition struct {

	// **参数解释**： 字段名称。 **取值范围**： systemQuery：是否隐藏系统查询。 userName：用户名称。 applicationName：应用名称。 dbName：数据库名称。 resourcePool：资源池。 queryStatus：查询状态。 enqueue：排队状态。
	Field string `json:"field"`

	// **参数解释**： 字段值。 **取值范围**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 比较方式。 **取值范围**： String类型参数：=、!=、like、not like int类型参数：=、!=、>、<、>=、<= boolean类型参数：=、!=
	Operator string `json:"operator"`
}

func (o ListQueriesCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueriesCondition struct{}"
	}

	return strings.Join([]string{"ListQueriesCondition", string(data)}, " ")
}
