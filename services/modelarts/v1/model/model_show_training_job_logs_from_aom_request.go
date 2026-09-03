package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowTrainingJobLogsFromAomRequest Request Object
type ShowTrainingJobLogsFromAomRequest struct {

	// 训练作业ID。获取方法请参见[查询训练作业列表](ListTrainingJobs.xml)。
	TrainingJobId string `json:"training_job_id"`

	// 训练作业的任务名称。可从训练作业详情中的status.tasks字段中获取。
	TaskId string `json:"task_id"`

	// **参数解释**：日志查询的基线行号，用于分页查询。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及（从最新日志开始查询）。
	BaseLine *string `json:"base_line,omitempty"`

	// **参数解释**：返回的日志行数。 **约束限制**：不涉及。 **取值范围**：1 ~ 500。 **默认取值**：50。
	Lines *int32 `json:"lines,omitempty"`

	// **参数解释**：日志排序方式。 **约束限制**：不涉及。 **取值范围**：枚举值如下： - asc：升序（从旧到新） - desc：降序（从新到旧） **默认取值**：desc。
	Order *ShowTrainingJobLogsFromAomRequestOrder `json:"order,omitempty"`
}

func (o ShowTrainingJobLogsFromAomRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTrainingJobLogsFromAomRequest struct{}"
	}

	return strings.Join([]string{"ShowTrainingJobLogsFromAomRequest", string(data)}, " ")
}

type ShowTrainingJobLogsFromAomRequestOrder struct {
	value string
}

type ShowTrainingJobLogsFromAomRequestOrderEnum struct {
	ASC  ShowTrainingJobLogsFromAomRequestOrder
	DESC ShowTrainingJobLogsFromAomRequestOrder
}

func GetShowTrainingJobLogsFromAomRequestOrderEnum() ShowTrainingJobLogsFromAomRequestOrderEnum {
	return ShowTrainingJobLogsFromAomRequestOrderEnum{
		ASC: ShowTrainingJobLogsFromAomRequestOrder{
			value: "asc",
		},
		DESC: ShowTrainingJobLogsFromAomRequestOrder{
			value: "desc",
		},
	}
}

func (c ShowTrainingJobLogsFromAomRequestOrder) Value() string {
	return c.value
}

func (c ShowTrainingJobLogsFromAomRequestOrder) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTrainingJobLogsFromAomRequestOrder) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
