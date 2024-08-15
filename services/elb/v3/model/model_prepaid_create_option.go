package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrepaidCreateOption 参数解释：创建负载均衡器实例的预付费计费配置。若传入该结构体，则创建预付费类型的负载均衡器实例。  [不支持该字段，请勿使用](tag:dt,dt_test,hcso_dt)
type PrepaidCreateOption struct {

	// 参数解释：预付费实例的订购周期类型，当前支持月和年。  取值范围：  - month：月。  - year：年。
	PeriodType PrepaidCreateOptionPeriodType `json:"period_type"`

	// 参数解释：预付费实例的订购周期数。  取值范围： - period_type为month时，为[1,9]。 - period_type为year时，为[1,3]。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 参数解释：自动续订开关。  取值范围： - true：开启自动续订。 - false：关闭自动续订。
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// 参数解释：自动支付开关。下单订购后，是否自动从客户的账户中支付。  约束限制：开启自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择关闭自动支付，然后在用户费用中心，选择代金券支付。  取值范围：  - true：开启自动支付。  - false：关闭自动支付。
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

func (c PrepaidCreateOptionPeriodType) Value() string {
	return c.value
}

func (c PrepaidCreateOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidCreateOptionPeriodType) UnmarshalJSON(b []byte) error {
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
