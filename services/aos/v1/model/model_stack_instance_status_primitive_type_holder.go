package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackInstanceStatusPrimitiveTypeHolder struct {

	// 资源栈实例的状态  * `WAIT_IN_PROGRESS` - 资源栈实例等待操作中 * `CANCEL_COMPLETE` - 资源栈实例操作取消完成 * `OPERATION_IN_PROGRESS` - 资源栈实例操作中 * `OPERATION_FAILED` - 资源栈实例操作失败 * `INOPERABLE` - 资源栈实例不可操作 * `OPERATION_COMPLETE` - 资源栈实例操作完成
	Status *StackInstanceStatusPrimitiveTypeHolderStatus `json:"status,omitempty"`
}

func (o StackInstanceStatusPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackInstanceStatusPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackInstanceStatusPrimitiveTypeHolder", string(data)}, " ")
}

type StackInstanceStatusPrimitiveTypeHolderStatus struct {
	value string
}

type StackInstanceStatusPrimitiveTypeHolderStatusEnum struct {
	WAIT_IN_PROGRESS      StackInstanceStatusPrimitiveTypeHolderStatus
	CANCEL_COMPLETE       StackInstanceStatusPrimitiveTypeHolderStatus
	OPERATION_IN_PROGRESS StackInstanceStatusPrimitiveTypeHolderStatus
	OPERATION_FAILED      StackInstanceStatusPrimitiveTypeHolderStatus
	INOPERABLE            StackInstanceStatusPrimitiveTypeHolderStatus
	OPERATION_COMPLETE    StackInstanceStatusPrimitiveTypeHolderStatus
}

func GetStackInstanceStatusPrimitiveTypeHolderStatusEnum() StackInstanceStatusPrimitiveTypeHolderStatusEnum {
	return StackInstanceStatusPrimitiveTypeHolderStatusEnum{
		WAIT_IN_PROGRESS: StackInstanceStatusPrimitiveTypeHolderStatus{
			value: "WAIT_IN_PROGRESS",
		},
		CANCEL_COMPLETE: StackInstanceStatusPrimitiveTypeHolderStatus{
			value: "CANCEL_COMPLETE",
		},
		OPERATION_IN_PROGRESS: StackInstanceStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		OPERATION_FAILED: StackInstanceStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_FAILED",
		},
		INOPERABLE: StackInstanceStatusPrimitiveTypeHolderStatus{
			value: "INOPERABLE",
		},
		OPERATION_COMPLETE: StackInstanceStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_COMPLETE",
		},
	}
}

func (c StackInstanceStatusPrimitiveTypeHolderStatus) Value() string {
	return c.value
}

func (c StackInstanceStatusPrimitiveTypeHolderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackInstanceStatusPrimitiveTypeHolderStatus) UnmarshalJSON(b []byte) error {
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
