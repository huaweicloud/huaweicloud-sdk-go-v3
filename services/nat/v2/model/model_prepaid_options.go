package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type PrepaidOptions struct {

	// 订购周期类型，当前支持包月和包年: - month: 包月 - year: 包年
	PeriodType PrepaidOptionsPeriodType `json:"period_type"`

	// 订购周期数，取值会随运营策略变化。 period_type为month时，为[1,9]， period_type为year时，为[1,3]
	PeriodNum int32 `json:"period_num"`

	// 是否自动续订； true：自动续订 false：不自动续订
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 下单订购后，是否自动从客户的账户中支付; true：自动支付 false：不自动支付。
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
