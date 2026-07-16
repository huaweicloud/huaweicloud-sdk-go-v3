package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CheckTrainingExperimentResponse Response Object
type CheckTrainingExperimentResponse struct {

	// **参数解释**：是否重复。 **约束限制**：不涉及。 **取值范围**： - true：重复 - false：不重复  **默认取值**：不涉及。
	IsDuplicate    *bool `json:"is_duplicate,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o CheckTrainingExperimentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckTrainingExperimentResponse struct{}"
	}

	return strings.Join([]string{"CheckTrainingExperimentResponse", string(data)}, " ")
}
