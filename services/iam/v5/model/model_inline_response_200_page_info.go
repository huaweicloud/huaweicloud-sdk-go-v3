package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// InlineResponse200PageInfo **参数解释**： 分页信息。  **取值范围**： 不涉及。
type InlineResponse200PageInfo struct {

	// **参数解释**： 如果存在，则表示还有后续的条目未显示在当前返回体中。请使用该值作为下一次请求的分页标记参数以获得下一页信息。请反复调用该接口直至该字段不存在。  **取值范围**： 不涉及。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**： 本页返回条目数量。  **取值范围**： 不涉及。
	CurrentCount int32 `json:"current_count"`
}

func (o InlineResponse200PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InlineResponse200PageInfo struct{}"
	}

	return strings.Join([]string{"InlineResponse200PageInfo", string(data)}, " ")
}
