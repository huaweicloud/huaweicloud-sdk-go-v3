package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PrepaidChangeChargeModeOption **参数解释**：包周期选项。  **约束限制**：pay_type=prepaid或charge_mode为prepaid时填写。
type PrepaidChangeChargeModeOption struct {

	// **参数解释**：是否连同独享按带宽计费的弹性公网IP一起转包周期。 弹性公网IP转包周期之后可以单独解绑、绑定到其他实例和删除EIP。  **约束限制**：只有独享且按带宽计费的弹性公网IP才被允许转包周期。  **取值范围**： - true：连同EIP一起转包周期。 - false：仅转ELB，不转EIP。  **默认取值**：false
	IncludePublicip *bool `json:"include_publicip,omitempty"`

	// **参数解释**：需要一起按需转包的弹性公网IP的ID。  **约束限制**： - 若include_publicip为false，不能指定该字段。 - 若include_publicip为true，该字段为未指定时，表示所有绑定的v4 eip都需要一起转包周期。 - 若include_publicip为true，该字段列表非空，表示只将指定的v4 eip转包。 - 若include_publicip为true，该字段列表为空，表示不指定任一eip转包，与include_publicip为false等效。  **取值范围**：不涉及  **默认取值**：不涉及
	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	// **参数解释**：订购周期类型。  **约束限制**：不涉及  **取值范围**： - month：月 - year：年  **默认取值**：month
	PeriodType PrepaidChangeChargeModeOptionPeriodType `json:"period_type"`

	// **参数解释**：订购周期数，取值会随运营策略变化。  **约束限制**：不涉及  **取值范围**： - period_type为month时，为[1,9] - period_type为year时，为[1,3]  **默认取值**：1
	PeriodNum *int32 `json:"period_num,omitempty"`

	// **参数解释**：是否自动续订。  **约束限制**：不涉及  **取值范围**： - true：自动续订。 - false：不自动续订。  **默认取值**：false
	AutoRenew *bool `json:"auto_renew,omitempty"`

	// **参数解释**：下单订购后，是否自动从客户的账户中支付。  **约束限制**：自动支付时，只能使用账户的现金支付；如果要使用代金券，请选择不自动支付，然后在用户费用中心，选择代金券支付。  **取值范围**： - true：自动支付。 - false：不自动支付。  **默认取值**：false
	AutoPay *bool `json:"auto_pay,omitempty"`
}

func (o PrepaidChangeChargeModeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PrepaidChangeChargeModeOption struct{}"
	}

	return strings.Join([]string{"PrepaidChangeChargeModeOption", string(data)}, " ")
}

type PrepaidChangeChargeModeOptionPeriodType struct {
	value string
}

type PrepaidChangeChargeModeOptionPeriodTypeEnum struct {
	MONTH PrepaidChangeChargeModeOptionPeriodType
	YEAR  PrepaidChangeChargeModeOptionPeriodType
}

func GetPrepaidChangeChargeModeOptionPeriodTypeEnum() PrepaidChangeChargeModeOptionPeriodTypeEnum {
	return PrepaidChangeChargeModeOptionPeriodTypeEnum{
		MONTH: PrepaidChangeChargeModeOptionPeriodType{
			value: "month",
		},
		YEAR: PrepaidChangeChargeModeOptionPeriodType{
			value: "year",
		},
	}
}

func (c PrepaidChangeChargeModeOptionPeriodType) Value() string {
	return c.value
}

func (c PrepaidChangeChargeModeOptionPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PrepaidChangeChargeModeOptionPeriodType) UnmarshalJSON(b []byte) error {
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
