package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 扩展参数说明
type CreateInstanceExtendParam struct {
	// 计费模式，取值范围： - prePaid：预付费，即包年/包月。 - postPaid：后付费，即按需付费。 默认值为postPaid。

	ChargeMode *CreateInstanceExtendParamChargeMode `json:"charge_mode,omitempty"`
	// 订购周期类型，取值范围： - month：月。 - year：年。 “charge_mode”参数配置为“prePaid”时该参数有效且为必选值。

	PeriodType *CreateInstanceExtendParamPeriodType `json:"period_type,omitempty"`
	// 订购周期数，取值范围： - period_type=month（周期类型为月）时，取值为[1，9]。 - period_type=year（周期类型为年）时，取值为1。 “charge_mode”参数配置为“prePaid”时该参数有效且为必选值。

	PeriodNum *int32 `json:"period_num,omitempty"`
	// 是否自动续订，取值范围： - “true”：自动续订。 - “false”：不自动续订。 “charge_mode”参数配置为“prePaid”时该参数有效，不传该字段时默认为不自动续订。\"

	IsAutoRenew *CreateInstanceExtendParamIsAutoRenew `json:"is_auto_renew,omitempty"`
	// 下单订购后，是否自动从客户的账户的余额中支付，取值范围： - “true”：是（自动从客户账户的余额中支付）。 - “false”：否（需要客户手动支付）。 “charge_mode”参数配置为“prePaid”时该参数有效，不传该字段时默认为客户手动支付。\"

	IsAutoPay *CreateInstanceExtendParamIsAutoPay `json:"is_auto_pay,omitempty"`
}

func (o CreateInstanceExtendParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceExtendParam struct{}"
	}

	return strings.Join([]string{"CreateInstanceExtendParam", string(data)}, " ")
}

type CreateInstanceExtendParamChargeMode struct {
	value string
}

type CreateInstanceExtendParamChargeModeEnum struct {
	PRE_PAID  CreateInstanceExtendParamChargeMode
	POST_PAID CreateInstanceExtendParamChargeMode
}

func GetCreateInstanceExtendParamChargeModeEnum() CreateInstanceExtendParamChargeModeEnum {
	return CreateInstanceExtendParamChargeModeEnum{
		PRE_PAID: CreateInstanceExtendParamChargeMode{
			value: "prePaid",
		},
		POST_PAID: CreateInstanceExtendParamChargeMode{
			value: "postPaid",
		},
	}
}

func (c CreateInstanceExtendParamChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceExtendParamChargeMode) UnmarshalJSON(b []byte) error {
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

type CreateInstanceExtendParamPeriodType struct {
	value string
}

type CreateInstanceExtendParamPeriodTypeEnum struct {
	MONTH CreateInstanceExtendParamPeriodType
	YEAR  CreateInstanceExtendParamPeriodType
}

func GetCreateInstanceExtendParamPeriodTypeEnum() CreateInstanceExtendParamPeriodTypeEnum {
	return CreateInstanceExtendParamPeriodTypeEnum{
		MONTH: CreateInstanceExtendParamPeriodType{
			value: "month",
		},
		YEAR: CreateInstanceExtendParamPeriodType{
			value: "year",
		},
	}
}

func (c CreateInstanceExtendParamPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceExtendParamPeriodType) UnmarshalJSON(b []byte) error {
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

type CreateInstanceExtendParamIsAutoRenew struct {
	value string
}

type CreateInstanceExtendParamIsAutoRenewEnum struct {
	TRUE  CreateInstanceExtendParamIsAutoRenew
	FALSE CreateInstanceExtendParamIsAutoRenew
}

func GetCreateInstanceExtendParamIsAutoRenewEnum() CreateInstanceExtendParamIsAutoRenewEnum {
	return CreateInstanceExtendParamIsAutoRenewEnum{
		TRUE: CreateInstanceExtendParamIsAutoRenew{
			value: "true",
		},
		FALSE: CreateInstanceExtendParamIsAutoRenew{
			value: "false",
		},
	}
}

func (c CreateInstanceExtendParamIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceExtendParamIsAutoRenew) UnmarshalJSON(b []byte) error {
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

type CreateInstanceExtendParamIsAutoPay struct {
	value string
}

type CreateInstanceExtendParamIsAutoPayEnum struct {
	TRUE  CreateInstanceExtendParamIsAutoPay
	FALSE CreateInstanceExtendParamIsAutoPay
}

func GetCreateInstanceExtendParamIsAutoPayEnum() CreateInstanceExtendParamIsAutoPayEnum {
	return CreateInstanceExtendParamIsAutoPayEnum{
		TRUE: CreateInstanceExtendParamIsAutoPay{
			value: "true",
		},
		FALSE: CreateInstanceExtendParamIsAutoPay{
			value: "false",
		},
	}
}

func (c CreateInstanceExtendParamIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateInstanceExtendParamIsAutoPay) UnmarshalJSON(b []byte) error {
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
