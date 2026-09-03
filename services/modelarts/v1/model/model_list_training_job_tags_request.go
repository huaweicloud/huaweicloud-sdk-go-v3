package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobTagsRequest Request Object
type ListTrainingJobTagsRequest struct {

	// **参数解释**：每页返回的标签key数量。 **约束限制**：不涉及。 **取值范围**：1 - 1000。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页查询的偏移量。 **约束限制**：不涉及。 **取值范围**：大于等于0。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListTrainingJobTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTrainingJobTagsRequest", string(data)}, " ")
}
