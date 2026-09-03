package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobsByTagsResponse Response Object
type ListTrainingJobsByTagsResponse struct {

	// **参数解释**：符合条件的训练作业资源列表。 **取值范围**：不涉及。
	Resources *[]ResourceInstance `json:"resources,omitempty"`

	// **参数解释**：符合条件的训练作业总数。 **取值范围**：不涉及。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListTrainingJobsByTagsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobsByTagsResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingJobsByTagsResponse", string(data)}, " ")
}
