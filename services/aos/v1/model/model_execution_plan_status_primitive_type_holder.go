package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecutionPlanStatusPrimitiveTypeHolder struct {

	// 执行计划的状态    * `CREATION_IN_PROGRESS` - 正在创建，请等待    * `CREATION_FAILED` - 创建失败，请从status_message获取错误信息汇总    * `AVAILABLE` - 创建完成，可以调用ApplyExecutionPlan API进行执行    * `APPLY_IN_PROGRESS` - 执行中，可通过GetStackMetadata查询资源栈状态，通过ListStackEvents获取执行过程中产生的资源栈事件    * `APPLIED` - 已执行
	Status *ExecutionPlanStatusPrimitiveTypeHolderStatus `json:"status,omitempty"`
}

func (o ExecutionPlanStatusPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecutionPlanStatusPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"ExecutionPlanStatusPrimitiveTypeHolder", string(data)}, " ")
}

type ExecutionPlanStatusPrimitiveTypeHolderStatus struct {
	value string
}

type ExecutionPlanStatusPrimitiveTypeHolderStatusEnum struct {
	CREATION_IN_PROGRESS ExecutionPlanStatusPrimitiveTypeHolderStatus
	CREATION_FAILED      ExecutionPlanStatusPrimitiveTypeHolderStatus
	AVAILABLE            ExecutionPlanStatusPrimitiveTypeHolderStatus
	APPLY_IN_PROGRESS    ExecutionPlanStatusPrimitiveTypeHolderStatus
	APPLIED              ExecutionPlanStatusPrimitiveTypeHolderStatus
}

func GetExecutionPlanStatusPrimitiveTypeHolderStatusEnum() ExecutionPlanStatusPrimitiveTypeHolderStatusEnum {
	return ExecutionPlanStatusPrimitiveTypeHolderStatusEnum{
		CREATION_IN_PROGRESS: ExecutionPlanStatusPrimitiveTypeHolderStatus{
			value: "CREATION_IN_PROGRESS",
		},
		CREATION_FAILED: ExecutionPlanStatusPrimitiveTypeHolderStatus{
			value: "CREATION_FAILED",
		},
		AVAILABLE: ExecutionPlanStatusPrimitiveTypeHolderStatus{
			value: "AVAILABLE",
		},
		APPLY_IN_PROGRESS: ExecutionPlanStatusPrimitiveTypeHolderStatus{
			value: "APPLY_IN_PROGRESS",
		},
		APPLIED: ExecutionPlanStatusPrimitiveTypeHolderStatus{
			value: "APPLIED",
		},
	}
}

func (c ExecutionPlanStatusPrimitiveTypeHolderStatus) Value() string {
	return c.value
}

func (c ExecutionPlanStatusPrimitiveTypeHolderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecutionPlanStatusPrimitiveTypeHolderStatus) UnmarshalJSON(b []byte) error {
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
