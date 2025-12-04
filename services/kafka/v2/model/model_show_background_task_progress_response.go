package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBackgroundTaskProgressResponse Response Object
type ShowBackgroundTaskProgressResponse struct {

	// **参数解释**： 进度。 **取值范围**： 不涉及。
	ProgressPercentage *int32 `json:"progress_percentage,omitempty"`

	// **参数解释**： 剩余时间。 **取值范围**： 不涉及。
	RemainTime *int32 `json:"remain_time,omitempty"`

	// **参数解释**： 步骤列表。
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
