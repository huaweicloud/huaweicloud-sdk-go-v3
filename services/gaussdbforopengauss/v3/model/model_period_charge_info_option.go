package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeriodChargeInfoOption **参数解释**: 包周期相关信息。 **约束限制**: 不涉及。
type PeriodChargeInfoOption struct {

	// **参数解释**: 周期单位类型。 **约束限制**: 不涉及。 **取值范围**: - year：包年。 - month：包月。  **默认取值**: 不涉及。
	PeriodType PeriodChargeInfoOptionPeriodType `json:"period_type"`

	// **参数解释**: 周期单位数量。 **约束限制**: 不涉及。 **取值范围**: - 当“period_type”为“month”时，取值为1~9。 - 当“period_type”为“year”时，取值为1~3。  当传入浮点型时，会自动截取为整型。 **默认取值**: 不涉及。
	PeriodNum int32 `json:"period_num"`

	// **参数解释**: 是否自动续费。 **约束限制**: 不涉及。 **取值范围**: - true：自动续订。 - false：不自动续订。  **默认取值**: false
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// **参数解释**: 是否自动支付。 **约束限制**: 不涉及。 **取值范围**: - true：自动支付。 - false：手动支付。  **默认取值**: false
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o PeriodChargeInfoOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodChargeInfoOption struct{}"
	}

	return strings.Join([]string{"PeriodChargeInfoOption", string(data)}, " ")
}

type PeriodChargeInfoOptionPeriodType struct {
	value string
}

type PeriodChargeInfoOptionPeriodTypeEnum struct {
	YEAR  PeriodChargeInfoOptionPeriodType
	MONTH PeriodChargeInfoOptionPeriodType
}

func GetPeriodChargeInfoOptionPeriodTypeEnum() PeriodChargeInfoOptionPeriodTypeEnum {
	return PeriodChargeInfoOptionPeriodTypeEnum{
		YEAR: PeriodChargeInfoOptionPeriodType{
			value: "year",
		},
		MONTH: PeriodChargeInfoOptionPeriodType{
			value: "month",
		},
	}
}

func (c PeriodChargeInfoOptionPeriodType) Value() string {
	return c.value
}

func (c PeriodChargeInfoOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodChargeInfoOptionPeriodType) UnmarshalJSON(b []byte) error {
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
