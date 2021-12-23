package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建负载均衡器的包周期信息，若传入该结构体，则创建包周期的LB
type PrepaidCreateOption struct {
	// 订购周期类型，当前支持包月和包年： month：月； year：年；

	PeriodType PrepaidCreateOptionPeriodType `json:"period_type"`
	// 订购周期数，取值会随运营策略变化。 period_type为month时，为[1,9]， period_type为year时，为[1,3]

	PeriodNum *int32 `json:"period_num,omitempty"`
	// 是否自动续订； true：自动续订 false：不自动续订

	AutoRenew *bool `json:"auto_renew,omitempty"`
	// 下单订购后，是否自动从客户的账户中支付； true：自动支付； false：不自动支付。 自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。

	AutoPay *bool `json:"auto_pay,omitempty"`
}

func (o PrepaidCreateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidCreateOption struct{}"
	}

	return strings.Join([]string{"PrepaidCreateOption", string(data)}, " ")
}

type PrepaidCreateOptionPeriodType struct {
	value string
}

type PrepaidCreateOptionPeriodTypeEnum struct {
	MONTH PrepaidCreateOptionPeriodType
	YEAR  PrepaidCreateOptionPeriodType
}

func GetPrepaidCreateOptionPeriodTypeEnum() PrepaidCreateOptionPeriodTypeEnum {
	return PrepaidCreateOptionPeriodTypeEnum{
		MONTH: PrepaidCreateOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidCreateOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidCreateOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidCreateOptionPeriodType) UnmarshalJSON(b []byte) error {
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
