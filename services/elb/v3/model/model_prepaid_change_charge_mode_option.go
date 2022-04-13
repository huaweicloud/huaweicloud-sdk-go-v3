package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 包周期选项，charge_mode为prepaid时填写。
type PrepaidChangeChargeModeOption struct {
	// 是否连同独享按带宽计费的弹性公网IP一起转包周期。 1. 弹性公网IP转包周期之后可以单独解绑，绑定到其他实例，删除 2. 只有独享且按带宽计费的弹性公网IP才被允许转包周期 默认值：false

	IncludePublicip *bool `json:"include_publicip,omitempty"`
	// 订购周期类型，仅在change_mode为delay时有效。取值： - month：月（默认）。 - year：年。

	PeriodType PrepaidChangeChargeModeOptionPeriodType `json:"period_type"`
	// 订购周期数，仅在change_mode为delay时有效。取值： - period_type为month时，为[1,9]，默认1。 - period_type为year时，为[1,3]，默认1。

	PeriodNum *int32 `json:"period_num,omitempty"`
	// 是否自动续订。取值： - true：自动续订 - false：不自动续订（默认）

	AutoRenew *bool `json:"auto_renew,omitempty"`
	// 下单订购后，是否自动从客户的账户中支付。取值： - true：自动支付； - false：不自动支付（默认）。  自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。

	AutoPay *bool `json:"auto_pay,omitempty"`
}

func (o PrepaidChangeChargeModeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidChangeChargeModeOption struct{}"
	}

	return strings.Join([]string{"PrepaidChangeChargeModeOption", string(data)}, " ")
}

type PrepaidChangeChargeModeOptionPeriodType struct {
	value string
}

type PrepaidChangeChargeModeOptionPeriodTypeEnum struct {
	MONTH PrepaidChangeChargeModeOptionPeriodType
	YEAR  PrepaidChangeChargeModeOptionPeriodType
}

func GetPrepaidChangeChargeModeOptionPeriodTypeEnum() PrepaidChangeChargeModeOptionPeriodTypeEnum {
	return PrepaidChangeChargeModeOptionPeriodTypeEnum{
		MONTH: PrepaidChangeChargeModeOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidChangeChargeModeOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidChangeChargeModeOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidChangeChargeModeOptionPeriodType) UnmarshalJSON(b []byte) error {
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
