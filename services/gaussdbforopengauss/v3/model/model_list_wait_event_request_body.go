package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListWaitEventRequestBody struct {

	// **参数解释**: 节点ID，仅支持包含有CN或DN（主、备）组件的节点。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	NodeId string `json:"node_id"`

	// **参数解释**: 是否查询系统用户。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	System *bool `json:"system,omitempty"`

	// **参数解释**: 查询记录数。 **约束限制**: 不能为负数。 **取值范围**: [1，100]。 **默认取值**: 默认为10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**: 索引位置，偏移量。从第一条数据偏移offset条数据后开始查询。 **约束限制**: 必须为数字，不能为负数。 **取值范围**: 不涉及。 **默认取值**: 默认为0（偏移0条数据，表示从第一条数据开始查询）。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**: 排序字段列表，内部List<String>。 **约束限制**: 不涉及。
	OrderFields *[][]string `json:"order_fields,omitempty"`

	WaitEventQueryInfo *WaitEventQueryInfo `json:"wait_event_query_info,omitempty"`
}

func (o ListWaitEventRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWaitEventRequestBody struct{}"
	}

	return strings.Join([]string{"ListWaitEventRequestBody", string(data)}, " ")
}
