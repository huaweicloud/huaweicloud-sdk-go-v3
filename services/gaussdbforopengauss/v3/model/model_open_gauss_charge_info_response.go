package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 计费类型信息，支持按需和包周期。
type OpenGaussChargeInfoResponse struct {

	// 计费模式。postPaid：后付费，即按需付费。prePaid：预付费，即包年/包月。
	ChargeMode OpenGaussChargeInfoResponseChargeMode `json:"charge_mode"`

	// 订购周期类型。month：包月。year：包年。 说明： “charge_mode”为“prePaid”时生效，且为必选值。
	PeriodType *OpenGaussChargeInfoResponsePeriodType `json:"period_type,omitempty"`

	// “charge_mode”为“prePaid”时生效，且为必选值，指定订购的时间。  取值范围：  当“period_type”为“month”时，取值为1~9。 当“period_type”为“year”时，取值为1~3。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 创建包周期实例时可指定，表示是否自动续订，续订的周期和原周期相同，且续订时会自动支付。  true，表示自动续订。 false，表示不自动续订，默认为该方式。
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// 创建包周期实例时可指定，表示是否自动从账户中支付，该字段不影响自动续订的支付方式。  true，表示自动从账户中支付。 false，表示手动从账户中支付，默认为该支付方式。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o OpenGaussChargeInfoResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenGaussChargeInfoResponse struct{}"
	}

	return strings.Join([]string{"OpenGaussChargeInfoResponse", string(data)}, " ")
}

type OpenGaussChargeInfoResponseChargeMode struct {
	value string
}

type OpenGaussChargeInfoResponseChargeModeEnum struct {
	POST_PAID OpenGaussChargeInfoResponseChargeMode
	PRE_PAID  OpenGaussChargeInfoResponseChargeMode
}

func GetOpenGaussChargeInfoResponseChargeModeEnum() OpenGaussChargeInfoResponseChargeModeEnum {
	return OpenGaussChargeInfoResponseChargeModeEnum{
		POST_PAID: OpenGaussChargeInfoResponseChargeMode{
			value: "postPaid",
		},
		PRE_PAID: OpenGaussChargeInfoResponseChargeMode{
			value: "prePaid",
		},
	}
}

func (c OpenGaussChargeInfoResponseChargeMode) Value() string {
	return c.value
}

func (c OpenGaussChargeInfoResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoResponseChargeMode) UnmarshalJSON(b []byte) error {
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

type OpenGaussChargeInfoResponsePeriodType struct {
	value string
}

type OpenGaussChargeInfoResponsePeriodTypeEnum struct {
	MONTH OpenGaussChargeInfoResponsePeriodType
	YEAR  OpenGaussChargeInfoResponsePeriodType
}

func GetOpenGaussChargeInfoResponsePeriodTypeEnum() OpenGaussChargeInfoResponsePeriodTypeEnum {
	return OpenGaussChargeInfoResponsePeriodTypeEnum{
		MONTH: OpenGaussChargeInfoResponsePeriodType{
			value: "month",
		},
		YEAR: OpenGaussChargeInfoResponsePeriodType{
			value: "year",
		},
	}
}

func (c OpenGaussChargeInfoResponsePeriodType) Value() string {
	return c.value
}

func (c OpenGaussChargeInfoResponsePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *OpenGaussChargeInfoResponsePeriodType) UnmarshalJSON(b []byte) error {
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
