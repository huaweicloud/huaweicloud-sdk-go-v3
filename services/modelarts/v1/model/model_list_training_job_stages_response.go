package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTrainingJobStagesResponse Response Object
type ListTrainingJobStagesResponse struct {

	// **参数解释**：总条数。 **取值范围**：不涉及。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：阶段记录。
	RunningRecords *[]StageRecord `json:"running_records,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListTrainingJobStagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTrainingJobStagesResponse struct{}"
	}

	return strings.Join([]string{"ListTrainingJobStagesResponse", string(data)}, " ")
}
