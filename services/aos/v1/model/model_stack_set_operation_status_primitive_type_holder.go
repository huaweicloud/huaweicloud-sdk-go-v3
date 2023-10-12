package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackSetOperationStatusPrimitiveTypeHolder struct {

	// 资源栈集操作状态   * `QUEUE_IN_PROGRESS` - 正在排队   * `OPERATION_IN_PROGRESS` - 正在操作   * `OPERATION_COMPLETE` - 操作完成   * `OPERATION_FAILED` - 操作失败   * `STOP_IN_PROGRESS` - 正在停止   * `STOP_COMPLETE` - 停止完成   * `STOP_FAILED` - 停止失败
	Status *StackSetOperationStatusPrimitiveTypeHolderStatus `json:"status,omitempty"`
}

func (o StackSetOperationStatusPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetOperationStatusPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetOperationStatusPrimitiveTypeHolder", string(data)}, " ")
}

type StackSetOperationStatusPrimitiveTypeHolderStatus struct {
	value string
}

type StackSetOperationStatusPrimitiveTypeHolderStatusEnum struct {
	QUEUE_IN_PROGRESS     StackSetOperationStatusPrimitiveTypeHolderStatus
	OPERATION_IN_PROGRESS StackSetOperationStatusPrimitiveTypeHolderStatus
	OPERATION_COMPLETE    StackSetOperationStatusPrimitiveTypeHolderStatus
	OPERATION_FAILED      StackSetOperationStatusPrimitiveTypeHolderStatus
	STOP_IN_PROGRESS      StackSetOperationStatusPrimitiveTypeHolderStatus
	STOP_COMPLETE         StackSetOperationStatusPrimitiveTypeHolderStatus
	STOP_FAILED           StackSetOperationStatusPrimitiveTypeHolderStatus
}

func GetStackSetOperationStatusPrimitiveTypeHolderStatusEnum() StackSetOperationStatusPrimitiveTypeHolderStatusEnum {
	return StackSetOperationStatusPrimitiveTypeHolderStatusEnum{
		QUEUE_IN_PROGRESS: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "QUEUE_IN_PROGRESS",
		},
		OPERATION_IN_PROGRESS: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		OPERATION_COMPLETE: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_COMPLETE",
		},
		OPERATION_FAILED: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_FAILED",
		},
		STOP_IN_PROGRESS: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "STOP_IN_PROGRESS",
		},
		STOP_COMPLETE: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "STOP_COMPLETE",
		},
		STOP_FAILED: StackSetOperationStatusPrimitiveTypeHolderStatus{
			value: "STOP_FAILED",
		},
	}
}

func (c StackSetOperationStatusPrimitiveTypeHolderStatus) Value() string {
	return c.value
}

func (c StackSetOperationStatusPrimitiveTypeHolderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetOperationStatusPrimitiveTypeHolderStatus) UnmarshalJSON(b []byte) error {
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
