package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPoolTagsRequest Request Object
type ListPoolTagsRequest struct {

	// **参数解释**：指定每一页查询返回的最大条目数，默认为200。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页列表的起始页，默认为0。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPoolTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPoolTagsRequest struct{}"
	}

	return strings.Join([]string{"ListPoolTagsRequest", string(data)}, " ")
}
