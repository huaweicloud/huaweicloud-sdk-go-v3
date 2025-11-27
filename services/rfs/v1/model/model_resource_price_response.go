package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ResourcePriceResponse struct {

	// 计费模式  * `PRE_PAID` - 包周期计费 * `POST_PAID` - 按需计费 * `FREE` - 免费
	ChargeMode *ResourcePriceResponseChargeMode `json:"charge_mode,omitempty"`

	// 该资源最终优惠后的金额（只考虑官网折扣、商务折扣以及伙伴折扣，不包含促销折扣及优惠券），保留小数点后2位，向上取整，默认单位是元。
	SalePrice *float64 `json:"sale_price,omitempty"`

	// 该资源的总优惠额，保留小数点后2位，向上取整，默认单位是元。
	Discount *float64 `json:"discount,omitempty"`

	// 该资源的原价，保留小数点后2位，向上取整，默认单位是元。
	OriginalPrice *float64 `json:"original_price,omitempty"`

	// 计费单位  如果该资源支持包周期计费或按需计费，则会返回该字段；如果该资源为免费资源，则不返回该字段。  * `HOUR` - 小时，按需计费的单位 * `DAY` - 天，按需计费的单位 * `MONTH` - 月，包周期计费的单位 * `YEAR` - 年，包周期计费的单位 * `BYTE` - 字节，按需计费的单位 * `MB` - 百万字节，包周期计费和按需计费的单位 * `GB` - 千兆字节，包周期计费和按需计费的单位
	PeriodType *ResourcePriceResponsePeriodType `json:"period_type,omitempty"`

	// 该资源的计费数量，需要和period_type搭配使用  如果该资源支持包周期计费或按需计费，则会返回该字段；如果该资源为免费资源，则不返回该字段。  * 对于按需计费资源，此值默认返回1，代表在1个计费单位下，该资源的价格 * 对于包周期计费资源，此值与模板中该资源的period字段保持一致
	PeriodCount *int32 `json:"period_count,omitempty"`

	// 该资源的最优折扣类型  对于按需计费资源的折扣类型 合同商务优惠类型：605 （华为云商务-折扣率，一口价，华为云用户）、606 （渠道商务-折扣率，一口价，BP用户） 伙伴折扣优惠类型：607 (合作伙伴授予折扣-折扣率)  对于包周期计费资源的折扣类型 合同商务折扣：605（华为云BE场景下的合同商务折扣）、606（分销商BE场景下的合同商务折扣） 伙伴授予折扣：607 促销折扣：700  如果该资源存在折扣，则返回该字段；如果该资源不存在折扣，则不返回该字段。
	BestDiscountType *int32 `json:"best_discount_type,omitempty"`

	// 最优折扣优惠额，保留小数点后2位，向上取整，默认单位是元。  如果该资源存在折扣，则返回该字段；如果该资源不存在折扣，则不返回该字段。
	BestDiscountPrice *float64 `json:"best_discount_price,omitempty"`

	// 官网价优惠额，保留小数点后2位，向上取整，默认单位是元。  如果该资源存在官网价优惠额，则返回该字段；如果该资源不存在官网价优惠额，则不返回该字段。
	OfficialWebsiteDiscountPrice *float64 `json:"official_website_discount_price,omitempty"`
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

type ResourcePriceResponsePeriodType struct {
	value string
}

type ResourcePriceResponsePeriodTypeEnum struct {
	HOUR  ResourcePriceResponsePeriodType
	DAY   ResourcePriceResponsePeriodType
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
