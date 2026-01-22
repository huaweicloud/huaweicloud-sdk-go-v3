package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackgroundTaskProgressResponse Response Object
type ShowBackgroundTaskProgressResponse struct {

	// **参数解释**： 后台任务处理进度。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ProgressPercentage *int32 `json:"progress_percentage,omitempty"`

	// **参数解释**： 剩余时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RemainTime *int32 `json:"remain_time,omitempty"`

	// **参数解释**： 步骤列表。   **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	StepList       *[]StepDetail `json:"step_list,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowBackgroundTaskProgressResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBackgroundTaskProgressResponse struct{}"
	}

	return strings.Join([]string{"ShowBackgroundTaskProgressResponse", string(data)}, " ")
}
