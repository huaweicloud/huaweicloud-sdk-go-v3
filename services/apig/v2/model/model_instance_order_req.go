package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceOrderReq struct {

	// 产品编号
	ProductId *string `json:"product_id,omitempty"`

	// 计费模式： - 0：按需 - 1：包周期
	ChargingMode *InstanceOrderReqChargingMode `json:"charging_mode,omitempty"`

	// 支付模式： - ALL_UPFRONT：全预付
	PaymentMode *InstanceOrderReqPaymentMode `json:"payment_mode,omitempty"`

	// 订购周期类型： - 2：月 - 3：年
	PeriodType *int32 `json:"period_type,omitempty"`

	// 订购周期数：1-9
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否支持自动续费： - 0：不自动续费 - 1：自动续费
	IsAutoRenew *InstanceOrderReqIsAutoRenew `json:"is_auto_renew,omitempty"`

	// 促销产品编号
	PromotionId *string `json:"promotion_id,omitempty"`

	// 促销计划编号
	PromotionPlanId *string `json:"promotion_plan_id,omitempty"`

	// 促销信息
	PromotionInfo *string `json:"promotion_info,omitempty"`

	// 组合产品编号
	CompositeProductId *string `json:"composite_product_id,omitempty"`

	InstanceInfo *InstanceCreateReqV2 `json:"instance_info,omitempty"`
}

func (o InstanceOrderReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceOrderReq struct{}"
	}

	return strings.Join([]string{"InstanceOrderReq", string(data)}, " ")
}

type InstanceOrderReqChargingMode struct {
	value int32
}

type InstanceOrderReqChargingModeEnum struct {
	E_0 InstanceOrderReqChargingMode
	E_1 InstanceOrderReqChargingMode
}

func GetInstanceOrderReqChargingModeEnum() InstanceOrderReqChargingModeEnum {
	return InstanceOrderReqChargingModeEnum{
		E_0: InstanceOrderReqChargingMode{
			value: 0,
		}, E_1: InstanceOrderReqChargingMode{
			value: 1,
		},
	}
}

func (c InstanceOrderReqChargingMode) Value() int32 {
	return c.value
}

func (c InstanceOrderReqChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceOrderReqChargingMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type InstanceOrderReqPaymentMode struct {
	value string
}

type InstanceOrderReqPaymentModeEnum struct {
	ALL_UPFRONT InstanceOrderReqPaymentMode
}

func GetInstanceOrderReqPaymentModeEnum() InstanceOrderReqPaymentModeEnum {
	return InstanceOrderReqPaymentModeEnum{
		ALL_UPFRONT: InstanceOrderReqPaymentMode{
			value: "ALL_UPFRONT",
		},
	}
}

func (c InstanceOrderReqPaymentMode) Value() string {
	return c.value
}

func (c InstanceOrderReqPaymentMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceOrderReqPaymentMode) UnmarshalJSON(b []byte) error {
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

type InstanceOrderReqIsAutoRenew struct {
	value int32
}

type InstanceOrderReqIsAutoRenewEnum struct {
	E_0 InstanceOrderReqIsAutoRenew
	E_1 InstanceOrderReqIsAutoRenew
}

func GetInstanceOrderReqIsAutoRenewEnum() InstanceOrderReqIsAutoRenewEnum {
	return InstanceOrderReqIsAutoRenewEnum{
		E_0: InstanceOrderReqIsAutoRenew{
			value: 0,
		}, E_1: InstanceOrderReqIsAutoRenew{
			value: 1,
		},
	}
}

func (c InstanceOrderReqIsAutoRenew) Value() int32 {
	return c.value
}

func (c InstanceOrderReqIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceOrderReqIsAutoRenew) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
