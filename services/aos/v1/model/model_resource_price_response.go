package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// resource price response
type ResourcePriceResponse struct {

	// 计费模式，包周期计费、按需计费、免费的返回，不支持的资源不返回 PRE_PAID 包周期计费 POST_PAID 按需计费 FREE 免费
	ChargeMode *ResourcePriceResponseChargeMode `json:"charge_mode,omitempty"`

	// 执行计划中的每个资源部署后最终优惠后的金额，保留小数点后2位，向下取整，默认单位是元
	SalePrice *interface{} `json:"sale_price,omitempty"`

	// 执行计划中的每个资源部署后的优惠额，保留小数点后2位，向下取整，默认单位是元
	Discount *interface{} `json:"discount,omitempty"`

	// 执行计划中的每个资源部署后的原价，保留小数点后2位，向下取整，默认单位是元
	OriginalPrice *interface{} `json:"original_price,omitempty"`

	// 包周期和按需的计费单位，包周期计费和按需计费返回，免费不会返回 HOUR：小时，包周期计费和按需计费返回，免费不会返回 DAY：天，包周期计费返回，按需计费和免费不会返回 WEEK：周，包周期计费返回，按需计费和免费不会返回 MONTH：月，包周期计费返回，按需计费和免费不会返回 YEAR：年，包周期计费返回，按需计费和免费不会返回 BYTE：字节，按需计费返回，包周期计费和免费不会返回 MB：百万字节，按需计费返回，包周期计费和免费不会返回 GB：千兆字节，按需计费返回，包周期计费和免费不会返回
	PeriodType *ResourcePriceResponsePeriodType `json:"period_type,omitempty"`

	// 订购数量。包周期计费和按需计费返回，免费不会返回。
	PeriodCount *int32 `json:"period_count,omitempty"`
}

func (o ResourcePriceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResourcePriceResponse struct{}"
	}

	return strings.Join([]string{"ResourcePriceResponse", string(data)}, " ")
}

type ResourcePriceResponseChargeMode struct {
	value string
}

type ResourcePriceResponseChargeModeEnum struct {
	PRE_PAID  ResourcePriceResponseChargeMode
	POST_PAID ResourcePriceResponseChargeMode
	FREE      ResourcePriceResponseChargeMode
}

func GetResourcePriceResponseChargeModeEnum() ResourcePriceResponseChargeModeEnum {
	return ResourcePriceResponseChargeModeEnum{
		PRE_PAID: ResourcePriceResponseChargeMode{
			value: "PRE_PAID",
		},
		POST_PAID: ResourcePriceResponseChargeMode{
			value: "POST_PAID",
		},
		FREE: ResourcePriceResponseChargeMode{
			value: "FREE",
		},
	}
}

func (c ResourcePriceResponseChargeMode) Value() string {
	return c.value
}

func (c ResourcePriceResponseChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcePriceResponseChargeMode) UnmarshalJSON(b []byte) error {
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

type ResourcePriceResponsePeriodType struct {
	value string
}

type ResourcePriceResponsePeriodTypeEnum struct {
	HOUR  ResourcePriceResponsePeriodType
	DAY   ResourcePriceResponsePeriodType
	WEEK  ResourcePriceResponsePeriodType
	MONTH ResourcePriceResponsePeriodType
	YEAR  ResourcePriceResponsePeriodType
	BYTE  ResourcePriceResponsePeriodType
	MB    ResourcePriceResponsePeriodType
	GB    ResourcePriceResponsePeriodType
}

func GetResourcePriceResponsePeriodTypeEnum() ResourcePriceResponsePeriodTypeEnum {
	return ResourcePriceResponsePeriodTypeEnum{
		HOUR: ResourcePriceResponsePeriodType{
			value: "HOUR",
		},
		DAY: ResourcePriceResponsePeriodType{
			value: "DAY",
		},
		WEEK: ResourcePriceResponsePeriodType{
			value: "WEEK",
		},
		MONTH: ResourcePriceResponsePeriodType{
			value: "MONTH",
		},
		YEAR: ResourcePriceResponsePeriodType{
			value: "YEAR",
		},
		BYTE: ResourcePriceResponsePeriodType{
			value: "BYTE",
		},
		MB: ResourcePriceResponsePeriodType{
			value: "MB",
		},
		GB: ResourcePriceResponsePeriodType{
			value: "GB",
		},
	}
}

func (c ResourcePriceResponsePeriodType) Value() string {
	return c.value
}

func (c ResourcePriceResponsePeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ResourcePriceResponsePeriodType) UnmarshalJSON(b []byte) error {
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
