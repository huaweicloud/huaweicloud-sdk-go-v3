package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolNodesRequest Request Object
type ListPoolNodesRequest struct {

	// **参数解释**：分页查询的偏移标志。 **约束限制**：可选。 **取值范围**：取值来自用户上一次分页查询响应结果中metadata.continue中的值，值为空默认无偏移。 **默认取值**：不涉及。
	Continue *string `json:"continue,omitempty"`

	// **参数解释**：分页单次查询返回的资源数量。 **约束限制**：不涉及。 **取值范围**：0 - 500。 **默认取值**：500。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：资源池的ID，取值自资源池详情的metadata.name字段。 **约束限制**：只能以小写字母开头，数字、中划线组成，不能以中划线结尾，且长度为36-63个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName string `json:"pool_name"`
}

func (o ListPoolNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolNodesRequest struct{}"
	}

	return strings.Join([]string{"ListPoolNodesRequest", string(data)}, " ")
}
