package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 状态。
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

func (c ExternalAccessStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExternalAccessStatus) UnmarshalJSON(b []byte) error {
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
