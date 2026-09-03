package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CountTrainingJobsByTagsResponse Response Object
type CountTrainingJobsByTagsResponse struct {

	// **参数解释**：符合条件的训练作业总数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CountTrainingJobsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountTrainingJobsByTagsResponse struct{}"
	}

	return strings.Join([]string{"CountTrainingJobsByTagsResponse", string(data)}, " ")
}
