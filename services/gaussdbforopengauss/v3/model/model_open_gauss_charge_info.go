package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// OpenGaussChargeInfo 计费类型信息，仅支持按需和包周期。
type OpenGaussChargeInfo struct {

	// 计费模式。postPaid：后付费，即按需付费。prePaid：预付费，即包年/包月。
	ChargeMode OpenGaussChargeInfoChargeMode `json:"charge_mode"`

	// 订购周期类型。month：包月。year：包年。 说明： “charge_mode”为“prePaid”时生效，且为必选值。
	PeriodType *OpenGaussChargeInfoPeriodType `json:"period_type,omitempty"`

	// “charge_mode”为“prePaid”时生效，且为必选值，指定订购的时间。  取值范围：  当“period_type”为“month”时，取值为1~9。 当“period_type”为“year”时，取值为1~3。  当传入浮点型时，会自动截取为整型。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 创建包周期实例时可指定，表示是否自动续订，续订时会自动支付。 按月订购时续订周期默认为1个月，按年订购时续订周期默认为1年，续订周期可自定义修改。  true，表示自动续订。 false，表示不自动续订，默认为该方式。
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 创建包周期实例时可指定，表示是否自动从账户中支付，该字段不影响自动续订的支付方式。  true，表示自动从账户中支付。 false，表示手动从账户中支付，默认为该支付方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o OpenGaussChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussChargeInfo struct{}"
	}

	return strings.Join([]string{"OpenGaussChargeInfo", string(data)}, " ")
}

type OpenGaussChargeInfoChargeMode struct {
	value string
}

type OpenGaussChargeInfoChargeModeEnum struct {
	POST_PAID OpenGaussChargeInfoChargeMode
	PRE_PAID  OpenGaussChargeInfoChargeMode
}

func GetOpenGaussChargeInfoChargeModeEnum() OpenGaussChargeInfoChargeModeEnum {
	return OpenGaussChargeInfoChargeModeEnum{
		POST_PAID: OpenGaussChargeInfoChargeMode{
			value: "postPaid",
		},
		PRE_PAID: OpenGaussChargeInfoChargeMode{
			value: "prePaid",
		},
	}
}

func (c OpenGaussChargeInfoChargeMode) Value() string {
	return c.value
}

func (c OpenGaussChargeInfoChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoChargeMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussChargeInfoPeriodType struct {
	value string
}

type OpenGaussChargeInfoPeriodTypeEnum struct {
	MONTH OpenGaussChargeInfoPeriodType
	YEAR  OpenGaussChargeInfoPeriodType
}

func GetOpenGaussChargeInfoPeriodTypeEnum() OpenGaussChargeInfoPeriodTypeEnum {
	return OpenGaussChargeInfoPeriodTypeEnum{
		MONTH: OpenGaussChargeInfoPeriodType{
			value: "month",
		},
		YEAR: OpenGaussChargeInfoPeriodType{
			value: "year",
		},
	}
}

func (c OpenGaussChargeInfoPeriodType) Value() string {
	return c.value
}

func (c OpenGaussChargeInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoPeriodType) UnmarshalJSON(b []byte) error {
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
