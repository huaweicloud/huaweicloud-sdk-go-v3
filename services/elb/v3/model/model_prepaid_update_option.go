package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrepaidUpdateOption 负载均衡器的包周期规格变更信息
type PrepaidUpdateOption struct {

	// 下单订购后，是否自动从客户的账户中支付； true：自动支付； false：不自动支付（默认）。 自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。
	AutoPay *bool `json:"auto_pay,omitempty"`

	// 规格变更类型： immediate：即时变更，规格变更立即生效。（默认） delay：续费变更，当前周期结束后变更为目标规格。
	ChangeMode *PrepaidUpdateOptionChangeMode `json:"change_mode,omitempty"`

	// 订购周期数（默认1），取值会随运营策略变化。（仅在change_mode为delay时生效） period_type为month时，为[1,9]， period_type为year时，为[1,3]
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 订购周期类型，当前支持包月和包年： （仅在change_mode为delay时生效） month：月（默认）； year：年；
	PeriodType *PrepaidUpdateOptionPeriodType `json:"period_type,omitempty"`
}

func (o PrepaidUpdateOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidUpdateOption struct{}"
	}

	return strings.Join([]string{"PrepaidUpdateOption", string(data)}, " ")
}

type PrepaidUpdateOptionChangeMode struct {
	value string
}

type PrepaidUpdateOptionChangeModeEnum struct {
	IMMEDIATE PrepaidUpdateOptionChangeMode
	DELAY     PrepaidUpdateOptionChangeMode
}

func GetPrepaidUpdateOptionChangeModeEnum() PrepaidUpdateOptionChangeModeEnum {
	return PrepaidUpdateOptionChangeModeEnum{
		IMMEDIATE: PrepaidUpdateOptionChangeMode{
			value: "immediate",
		},
		DELAY: PrepaidUpdateOptionChangeMode{
			value: "delay",
		},
	}
}

func (c PrepaidUpdateOptionChangeMode) Value() string {
	return c.value
}

func (c PrepaidUpdateOptionChangeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidUpdateOptionChangeMode) UnmarshalJSON(b []byte) error {
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

type PrepaidUpdateOptionPeriodType struct {
	value string
}

type PrepaidUpdateOptionPeriodTypeEnum struct {
	MONTH PrepaidUpdateOptionPeriodType
	YEAR  PrepaidUpdateOptionPeriodType
}

func GetPrepaidUpdateOptionPeriodTypeEnum() PrepaidUpdateOptionPeriodTypeEnum {
	return PrepaidUpdateOptionPeriodTypeEnum{
		MONTH: PrepaidUpdateOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidUpdateOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidUpdateOptionPeriodType) Value() string {
	return c.value
}

func (c PrepaidUpdateOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidUpdateOptionPeriodType) UnmarshalJSON(b []byte) error {
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
