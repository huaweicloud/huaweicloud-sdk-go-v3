package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeListMetadata 节点的metadata信息。
type NodeListMetadata struct {

	// **参数解释**：分页查询的偏移标志。 **取值范围**：取值来自用户上一次分页查询响应结果中metadata.continue中的值，值为空默认无偏移。
	Continue *string `json:"continue,omitempty"`

	// **参数解释**：分页查询中剩余资源的数量。 **默认取值**：不涉及。
	RemainingItemCount *int64 `json:"remainingItemCount,omitempty"`
}

func (o NodeListMetadata) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeListMetadata struct{}"
	}

	return strings.Join([]string{"NodeListMetadata", string(data)}, " ")
}
