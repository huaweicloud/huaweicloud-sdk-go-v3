package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobsByTagsRequest Request Object
type ListTrainingJobsByTagsRequest struct {

	// **参数解释**：每页返回的资源数量。 **约束限制**：不涉及。 **取值范围**：1 - 1000。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：分页查询的偏移量。 **约束限制**：不涉及。 **取值范围**：大于等于0。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	Body *ResourceInstancesFilterReq `json:"body,omitempty"`
}

func (o ListTrainingJobsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobsByTagsRequest struct{}"
	}

	return strings.Join([]string{"ListTrainingJobsByTagsRequest", string(data)}, " ")
}
