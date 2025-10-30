package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OperationStatus 运维状态说明： - FREEZE - 已冻结
type OperationStatus struct {
	value string
}

type OperationStatusEnum struct {
	FREEZE OperationStatus
}

func GetOperationStatusEnum() OperationStatusEnum {
	return OperationStatusEnum{
		FREEZE: OperationStatus{
			value: "FREEZE",
		},
	}
}

func (c OperationStatus) Value() string {
	return c.value
}

func (c OperationStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OperationStatus) UnmarshalJSON(b []byte) error {
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
