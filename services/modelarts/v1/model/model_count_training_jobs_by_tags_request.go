package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTrainingJobsByTagsRequest Request Object
type CountTrainingJobsByTagsRequest struct {

	// **参数解释**：返回的数据条目数。 **约束限制**：不涉及。 **取值范围**：1 ~ 1000。 **默认取值**：1000。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：偏移量，表示从第几条数据开始查询。 **约束限制**：不涉及。 **取值范围**：不小于0。 **默认取值**：0。
	Offset *int32 `json:"offset,omitempty"`

	Body *CountResourceInstancesReq `json:"body,omitempty"`
}

func (o CountTrainingJobsByTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTrainingJobsByTagsRequest struct{}"
	}

	return strings.Join([]string{"CountTrainingJobsByTagsRequest", string(data)}, " ")
}
