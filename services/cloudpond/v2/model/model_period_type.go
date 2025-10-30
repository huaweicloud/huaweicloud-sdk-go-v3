package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeriodType 包周期类型 - year：包年
type PeriodType struct {
	value string
}

type PeriodTypeEnum struct {
	YEAR PeriodType
}

func GetPeriodTypeEnum() PeriodTypeEnum {
	return PeriodTypeEnum{
		YEAR: PeriodType{
			value: "year",
		},
	}
}

func (c PeriodType) Value() string {
	return c.value
}

func (c PeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodType) UnmarshalJSON(b []byte) error {
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
