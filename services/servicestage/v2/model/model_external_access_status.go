package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ExternalAccessStatus 状态。
type ExternalAccessStatus struct {
	value string
}

type ExternalAccessStatusEnum struct {
	NORMAL  ExternalAccessStatus
	EXPIRED ExternalAccessStatus
}

func GetExternalAccessStatusEnum() ExternalAccessStatusEnum {
	return ExternalAccessStatusEnum{
		NORMAL: ExternalAccessStatus{
			value: "NORMAL",
		},
		EXPIRED: ExternalAccessStatus{
			value: "EXPIRED",
		},
	}
}

func (c ExternalAccessStatus) Value() string {
	return c.value
}

func (c ExternalAccessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExternalAccessStatus) UnmarshalJSON(b []byte) error {
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
