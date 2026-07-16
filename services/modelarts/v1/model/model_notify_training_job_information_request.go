package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotifyTrainingJobInformationRequest Request Object
type NotifyTrainingJobInformationRequest struct {

	// **参数解释**：训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TrainingJobId string `json:"training_job_id"`

	// **参数解释**：训练作业的任务名称。可从训练作业详情中的status.tasks字段中获取。 **约束限制**：单节点默认为\"worker-0\"，多节点则为\"worker-0\"、\"worker-1\"，依次类推。 **取值范围**：不涉及。 **默认取值**：不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**：事件上报类型。 **约束限制**：不涉及。 **取值范围**：取\"training-event\"。 **默认取值**：不涉及。
	ReportType string `json:"report_type"`

	Body *ReportEventBody `json:"body,omitempty"`
}

func (o NotifyTrainingJobInformationRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotifyTrainingJobInformationRequest struct{}"
	}

	return strings.Join([]string{"NotifyTrainingJobInformationRequest", string(data)}, " ")
}
