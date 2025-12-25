package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListJobDetailsResponse Response Object
type ListJobDetailsResponse struct {

	// **参数解释**： 任务ID。 **取值范围**： 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**： 任务名称。 **取值范围**： 不涉及。
	JobName *string `json:"job_name,omitempty"`

	// **参数解释**： 任务开始时间。 **取值范围**： 不涉及。
	BeginTime *string `json:"begin_time,omitempty"`

	// **参数解释**： 任务结束时间。 **取值范围**： 不涉及。
	EndTime *string `json:"end_time,omitempty"`

	// **参数解释**： 任务当前状态。 **取值范围**： - INIT：初始化。 - WAITING：等待中。 - RUNNING：运行中。 - DELAY_SCHEDULED：延迟调度。 - FAIL：失败。 - SUCCESS：成功。 - STOP：停止。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 任务失败错误码。 **取值范围**： 不涉及。
	FailedCode *string `json:"failed_code,omitempty"`

	// **参数解释**： 任务失败错误详情。 **取值范围**： 不涉及。
	FailedDetail *string `json:"failed_detail,omitempty"`

	// **参数解释**： 任务进度。 **取值范围**： 不涉及。
	Progress       *string `json:"progress,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListJobDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListJobDetailsResponse struct{}"
	}

	return strings.Join([]string{"ListJobDetailsResponse", string(data)}, " ")
}
