package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTransactionRequestBody struct {

	// **参数解释**: 节点ID，仅支持非日志类型的CN或DN节点。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 组件ID，仅支持非日志类型的CN或DN节点。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ComponentId string `json:"component_id"`

	// **参数解释**: 查询结果的事务最大个数。 **约束限制**: 不涉及。 **取值范围**: [1，100]。 **默认取值**: 默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 查询结果的事务起始页码。 **约束限制**: 不涉及。 **取值范围**: [0，2^31-1]。 **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	TransactionQueryInfo *ListTransactionRequestBodyTransactionQueryInfo `json:"transaction_query_info,omitempty"`
}

func (o ListTransactionRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTransactionRequestBody struct{}"
	}

	return strings.Join([]string{"ListTransactionRequestBody", string(data)}, " ")
}
