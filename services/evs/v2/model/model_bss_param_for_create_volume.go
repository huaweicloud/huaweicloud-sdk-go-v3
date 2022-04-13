package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 包周期创卷的计费策略参数。
type BssParamForCreateVolume struct {
	// 功能说明：计费模式。默认值为postPaid。 取值范围： * prePaid：包年包月 * postPaid：按需

	ChargingMode *BssParamForCreateVolumeChargingMode `json:"chargingMode,omitempty"`
	// 功能说明：是否立即支付。chargingMode为PrePaid时该参数会生效。默认值为false。 取值范围： * true：立即支付，从帐户余额中自动扣费 * false：不立即支付，创建订单暂不支付

	IsAutoPay *BssParamForCreateVolumeIsAutoPay `json:"isAutoPay,omitempty"`
	// 功能说明：是否自动续订。chargingMode为prePaid时该参数会生效。默认值为false。 取值范围： * true：自动续订，自动续订周期与订购周期相同 * false：不自动续订

	IsAutoRenew *BssParamForCreateVolumeIsAutoRenew `json:"isAutoRenew,omitempty"`
	// 功能说明：订购周期数，chargingMode为prePaid时该参数会生效，并且该参数为为必选。 取值范围： * periodType为month时，为[1-9] * periodType为year时，为[1-1]

	PeriodNum *int32 `json:"periodNum,omitempty"`
	// 功能说明：订购周期单位。chargingMode为prePaid时该参数会生效，并且该参数为必选。 取值范围： * month：月 * year：年

	PeriodType *BssParamForCreateVolumePeriodType `json:"periodType,omitempty"`
}

func (o BssParamForCreateVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssParamForCreateVolume struct{}"
	}

	return strings.Join([]string{"BssParamForCreateVolume", string(data)}, " ")
}

type BssParamForCreateVolumeChargingMode struct {
	value string
}

type BssParamForCreateVolumeChargingModeEnum struct {
	POST_PAID BssParamForCreateVolumeChargingMode
	PRE_PAID  BssParamForCreateVolumeChargingMode
}

func GetBssParamForCreateVolumeChargingModeEnum() BssParamForCreateVolumeChargingModeEnum {
	return BssParamForCreateVolumeChargingModeEnum{
		POST_PAID: BssParamForCreateVolumeChargingMode{
			value: "postPaid",
		},
		PRE_PAID: BssParamForCreateVolumeChargingMode{
			value: "prePaid",
		},
	}
}

func (c BssParamForCreateVolumeChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumeChargingMode) UnmarshalJSON(b []byte) error {
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

type BssParamForCreateVolumeIsAutoPay struct {
	value string
}

type BssParamForCreateVolumeIsAutoPayEnum struct {
	TRUE  BssParamForCreateVolumeIsAutoPay
	FALSE BssParamForCreateVolumeIsAutoPay
}

func GetBssParamForCreateVolumeIsAutoPayEnum() BssParamForCreateVolumeIsAutoPayEnum {
	return BssParamForCreateVolumeIsAutoPayEnum{
		TRUE: BssParamForCreateVolumeIsAutoPay{
			value: "true",
		},
		FALSE: BssParamForCreateVolumeIsAutoPay{
			value: "false",
		},
	}
}

func (c BssParamForCreateVolumeIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumeIsAutoPay) UnmarshalJSON(b []byte) error {
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

type BssParamForCreateVolumeIsAutoRenew struct {
	value string
}

type BssParamForCreateVolumeIsAutoRenewEnum struct {
	TRUE  BssParamForCreateVolumeIsAutoRenew
	FALSE BssParamForCreateVolumeIsAutoRenew
}

func GetBssParamForCreateVolumeIsAutoRenewEnum() BssParamForCreateVolumeIsAutoRenewEnum {
	return BssParamForCreateVolumeIsAutoRenewEnum{
		TRUE: BssParamForCreateVolumeIsAutoRenew{
			value: "true",
		},
		FALSE: BssParamForCreateVolumeIsAutoRenew{
			value: "false",
		},
	}
}

func (c BssParamForCreateVolumeIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumeIsAutoRenew) UnmarshalJSON(b []byte) error {
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

type BssParamForCreateVolumePeriodType struct {
	value string
}

type BssParamForCreateVolumePeriodTypeEnum struct {
	MONTH BssParamForCreateVolumePeriodType
	YEAR  BssParamForCreateVolumePeriodType
}

func GetBssParamForCreateVolumePeriodTypeEnum() BssParamForCreateVolumePeriodTypeEnum {
	return BssParamForCreateVolumePeriodTypeEnum{
		MONTH: BssParamForCreateVolumePeriodType{
			value: "month",
		},
		YEAR: BssParamForCreateVolumePeriodType{
			value: "year",
		},
	}
}

func (c BssParamForCreateVolumePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssParamForCreateVolumePeriodType) UnmarshalJSON(b []byte) error {
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
