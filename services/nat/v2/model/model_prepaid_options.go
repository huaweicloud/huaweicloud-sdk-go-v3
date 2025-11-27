package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PrepaidOptions struct {

	// month: 包月 year: 包年
	PeriodType *PrepaidOptionsPeriodType `json:"period_type,omitempty"`

	// 周期大小
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动续费
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 是否自动支付
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o PrepaidOptions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidOptions struct{}"
	}

	return strings.Join([]string{"PrepaidOptions", string(data)}, " ")
}

type PrepaidOptionsPeriodType struct {
	value string
}

type PrepaidOptionsPeriodTypeEnum struct {
	MONTH PrepaidOptionsPeriodType
	YEAR  PrepaidOptionsPeriodType
}

func GetPrepaidOptionsPeriodTypeEnum() PrepaidOptionsPeriodTypeEnum {
	return PrepaidOptionsPeriodTypeEnum{
		MONTH: PrepaidOptionsPeriodType{
			value: "month",
		},
		YEAR: PrepaidOptionsPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidOptionsPeriodType) Value() string {
	return c.value
}

func (c PrepaidOptionsPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidOptionsPeriodType) UnmarshalJSON(b []byte) error {
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
