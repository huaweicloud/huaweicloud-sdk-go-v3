package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 执行计划的元数据
type ExecutionPlan struct {

	// 栈的唯一Id,为uuid
	StackId *string `json:"stack_id,omitempty"`

	// 栈的名字
	StackName *string `json:"stack_name,omitempty"`

	// 执行计划的唯一Id，由IaC随机生成,为uuid
	ExecutionPlanId *string `json:"execution_plan_id,omitempty"`

	// 执行计划的名字
	ExecutionPlanName *string `json:"execution_plan_name,omitempty"`

	// 执行计划的描述，此描述为用户在生成时指定
	Description *string `json:"description,omitempty"`

	// 执行计划的生成时间
	CreateTime *string `json:"create_time,omitempty"`

	// 执行时间
	ApplyTime *string `json:"apply_time,omitempty"`

	// 执行计划的执行状态，只有当AVAILABLE的时候才可以使用apply执行     * `CREATION_IN_PROGRESS` - 正在生成     * `CREATION_FAILED` - 生成失败     * `AVAILABLE` - 执行计划已经生成完成。可以使用apply进行执行     * `APPLIED` - 执行完成
	Status *ExecutionPlanStatus `json:"status,omitempty"`

	// 展示执行计划状态更多细节的信息
	StatusMessage *string `json:"status_message,omitempty"`
}

func (o ExecutionPlan) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlan struct{}"
	}

	return strings.Join([]string{"ExecutionPlan", string(data)}, " ")
}

type ExecutionPlanStatus struct {
	value string
}

type ExecutionPlanStatusEnum struct {
	CREATION_IN_PROGRESS ExecutionPlanStatus
	CREATION_FAILED      ExecutionPlanStatus
	AVAILABLE            ExecutionPlanStatus
	APPLIED              ExecutionPlanStatus
}

func GetExecutionPlanStatusEnum() ExecutionPlanStatusEnum {
	return ExecutionPlanStatusEnum{
		CREATION_IN_PROGRESS: ExecutionPlanStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: ExecutionPlanStatus{
			value: "CREATION_FAILED",
		},
		AVAILABLE: ExecutionPlanStatus{
			value: "AVAILABLE",
		},
		APPLIED: ExecutionPlanStatus{
			value: "APPLIED",
		},
	}
}

func (c ExecutionPlanStatus) Value() string {
	return c.value
}

func (c ExecutionPlanStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionPlanStatus) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
