package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type StackSetStatusPrimitiveTypeHolder struct {

	// 资源栈集的状态     * `IDLE` - 资源栈集空闲 * `OPERATION_IN_PROGRESS` - 资源栈集操作中 * `DEACTIVATED` - 资源栈集禁用
	Status *StackSetStatusPrimitiveTypeHolderStatus `json:"status,omitempty"`
}

func (o StackSetStatusPrimitiveTypeHolder) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StackSetStatusPrimitiveTypeHolder struct{}"
	}

	return strings.Join([]string{"StackSetStatusPrimitiveTypeHolder", string(data)}, " ")
}

type StackSetStatusPrimitiveTypeHolderStatus struct {
	value string
}

type StackSetStatusPrimitiveTypeHolderStatusEnum struct {
	IDLE                  StackSetStatusPrimitiveTypeHolderStatus
	OPERATION_IN_PROGRESS StackSetStatusPrimitiveTypeHolderStatus
	DEACTIVATED           StackSetStatusPrimitiveTypeHolderStatus
}

func GetStackSetStatusPrimitiveTypeHolderStatusEnum() StackSetStatusPrimitiveTypeHolderStatusEnum {
	return StackSetStatusPrimitiveTypeHolderStatusEnum{
		IDLE: StackSetStatusPrimitiveTypeHolderStatus{
			value: "IDLE",
		},
		OPERATION_IN_PROGRESS: StackSetStatusPrimitiveTypeHolderStatus{
			value: "OPERATION_IN_PROGRESS",
		},
		DEACTIVATED: StackSetStatusPrimitiveTypeHolderStatus{
			value: "DEACTIVATED",
		},
	}
}

func (c StackSetStatusPrimitiveTypeHolderStatus) Value() string {
	return c.value
}

func (c StackSetStatusPrimitiveTypeHolderStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StackSetStatusPrimitiveTypeHolderStatus) UnmarshalJSON(b []byte) error {
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
