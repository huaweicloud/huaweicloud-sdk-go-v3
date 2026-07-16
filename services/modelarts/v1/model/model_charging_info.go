package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ChargingInfo struct {

	// **参数解释**：付费类型。表示服务器的计费模式。 **约束限制**：不涉及。 **取值范围**： - COMMON：同时支持包周期和按需 - POST_PAID：后付费 - PRE_PAID：预付费 **默认取值**：不涉及。
	ChargingMode ChargingInfoChargingMode `json:"charging_mode"`

	// **参数解释**：是否自动支付。表示是否开启自动支付功能。 **约束限制**：不涉及。 **取值范围**： - true：自动支付 - false：不自动支付 **默认取值**：不涉及。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`

	// **参数解释**：是否自动续订。表示是否开启自动续订功能。 **约束限制**：不涉及。 **取值范围**： - true：自动续订 - false：不自动续订 **默认取值**：不涉及。
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// **参数解释**：订购周期数量。表示订购周期的数量。 **约束限制**：不涉及。 **取值范围**：1 - 11 **默认取值**：不涉及。
	PeriodNum int32 `json:"period_num"`

	// **参数解释**：订购周期类型。表示订购周期的时间单位。 **约束限制**：不涉及。 **取值范围**： - ABSOLUTE - DAY：天 - HOUR：小时 - MONTH：月 - WEEK：周 - YEAR：年 **默认取值**：不涉及。
	PeriodType ChargingInfoPeriodType `json:"period_type"`
}

func (o ChargingInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChargingInfo struct{}"
	}

	return strings.Join([]string{"ChargingInfo", string(data)}, " ")
}

type ChargingInfoChargingMode struct {
	value string
}

type ChargingInfoChargingModeEnum struct {
	COMMON    ChargingInfoChargingMode
	POST_PAID ChargingInfoChargingMode
	PRE_PAID  ChargingInfoChargingMode
}

func GetChargingInfoChargingModeEnum() ChargingInfoChargingModeEnum {
	return ChargingInfoChargingModeEnum{
		COMMON: ChargingInfoChargingMode{
			value: "COMMON",
		},
		POST_PAID: ChargingInfoChargingMode{
			value: "POST_PAID",
		},
		PRE_PAID: ChargingInfoChargingMode{
			value: "PRE_PAID",
		},
	}
}

func (c ChargingInfoChargingMode) Value() string {
	return c.value
}

func (c ChargingInfoChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChargingInfoChargingMode) UnmarshalJSON(b []byte) error {
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

type ChargingInfoPeriodType struct {
	value string
}

type ChargingInfoPeriodTypeEnum struct {
	ABSOLUTE ChargingInfoPeriodType
	DAY      ChargingInfoPeriodType
	HOUR     ChargingInfoPeriodType
	MONTH    ChargingInfoPeriodType
	WEEK     ChargingInfoPeriodType
	YEAR     ChargingInfoPeriodType
}

func GetChargingInfoPeriodTypeEnum() ChargingInfoPeriodTypeEnum {
	return ChargingInfoPeriodTypeEnum{
		ABSOLUTE: ChargingInfoPeriodType{
			value: "ABSOLUTE",
		},
		DAY: ChargingInfoPeriodType{
			value: "DAY",
		},
		HOUR: ChargingInfoPeriodType{
			value: "HOUR",
		},
		MONTH: ChargingInfoPeriodType{
			value: "MONTH",
		},
		WEEK: ChargingInfoPeriodType{
			value: "WEEK",
		},
		YEAR: ChargingInfoPeriodType{
			value: "YEAR",
		},
	}
}

func (c ChargingInfoPeriodType) Value() string {
	return c.value
}

func (c ChargingInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChargingInfoPeriodType) UnmarshalJSON(b []byte) error {
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
