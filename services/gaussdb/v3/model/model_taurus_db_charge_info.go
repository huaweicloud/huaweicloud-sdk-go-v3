package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaurusDbChargeInfo **参数解释**：  计费类型信息。  **约束限制**：  不涉及。
type TaurusDbChargeInfo struct {

	// **参数解释**：  订购周期类型。  **约束限制**：  不涉及。  **取值范围**：  - month：包月。 - year：包年。  **默认取值**：  不涉及。
	PeriodType TaurusDbChargeInfoPeriodType `json:"period_type"`

	// **参数解释**：  订购时间长度。  **约束限制**：  需要和period_type允许的订购时间长度对应。  **取值范围**：  - \"period_type\"为\"month\"时，取值为1~9。 - \"period_type\"为\"year\"时，取值为1~3。  **默认取值**：  不涉及。
	PeriodNum int32 `json:"period_num"`

	// **参数解释**：  是否自动续订，续订的周期和原周期相同，且续订时会自动支付。  **约束限制**：  不涉及。  **取值范围**：  - true：自动续订。 - false：不自动续订。  **默认取值**：  false。
	IsAutoRenew *bool `json:"is_auto_renew,omitempty"`

	// **参数解释**：  是否自动从客户的账户中支付，此字段不影响自动续订的支付方式。  **约束限制**：  不涉及。  **取值范围**：  - true：自动支付。 - false：手动支付。  **默认取值**：  false。
	IsAutoPay *bool `json:"is_auto_pay,omitempty"`
}

func (o TaurusDbChargeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaurusDbChargeInfo struct{}"
	}

	return strings.Join([]string{"TaurusDbChargeInfo", string(data)}, " ")
}

type TaurusDbChargeInfoPeriodType struct {
	value string
}

type TaurusDbChargeInfoPeriodTypeEnum struct {
	MONTH TaurusDbChargeInfoPeriodType
	YEAR  TaurusDbChargeInfoPeriodType
}

func GetTaurusDbChargeInfoPeriodTypeEnum() TaurusDbChargeInfoPeriodTypeEnum {
	return TaurusDbChargeInfoPeriodTypeEnum{
		MONTH: TaurusDbChargeInfoPeriodType{
			value: "month",
		},
		YEAR: TaurusDbChargeInfoPeriodType{
			value: "year",
		},
	}
}

func (c TaurusDbChargeInfoPeriodType) Value() string {
	return c.value
}

func (c TaurusDbChargeInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaurusDbChargeInfoPeriodType) UnmarshalJSON(b []byte) error {
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
