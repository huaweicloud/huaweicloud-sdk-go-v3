package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListNetworksRequest Request Object
type ListNetworksRequest struct {

	// **参数解释**：标签筛选。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	LabelSelector *string `json:"labelSelector,omitempty"`

	// **参数解释**：分页单次查询返回的资源数量。 **约束限制**：不涉及。 **取值范围**：0 - 500。 **默认取值**：500。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页查询的偏移标志。取值来自用户上一次分页查询响应结果中metadata.continue中的值，值为空默认无偏移。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Continue *string `json:"continue,omitempty"`
}

func (o ListNetworksRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListNetworksRequest struct{}"
	}

	return strings.Join([]string{"ListNetworksRequest", string(data)}, " ")
}
