package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ChangeLoadbalancerChargeModeRequestBody This is a auto create Body Object
type ChangeLoadbalancerChargeModeRequestBody struct {

	// **参数解释**：需要修改计费类型的负载均衡器ID列表。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	LoadbalancerIds []string `json:"loadbalancer_ids"`

	// **参数解释**：计费模式。  **约束限制**：不涉及  **取值范围**：仅支持设置为'prepaid'-包周期计费。  **默认取值**：不涉及
	ChargeMode ChangeLoadbalancerChargeModeRequestBodyChargeMode `json:"charge_mode"`

	PrepaidOptions *PrepaidChangeChargeModeOption `json:"prepaid_options,omitempty"`
}

func (o ChangeLoadbalancerChargeModeRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeLoadbalancerChargeModeRequestBody struct{}"
	}

	return strings.Join([]string{"ChangeLoadbalancerChargeModeRequestBody", string(data)}, " ")
}

type ChangeLoadbalancerChargeModeRequestBodyChargeMode struct {
	value string
}

type ChangeLoadbalancerChargeModeRequestBodyChargeModeEnum struct {
	PREPAID ChangeLoadbalancerChargeModeRequestBodyChargeMode
}

func GetChangeLoadbalancerChargeModeRequestBodyChargeModeEnum() ChangeLoadbalancerChargeModeRequestBodyChargeModeEnum {
	return ChangeLoadbalancerChargeModeRequestBodyChargeModeEnum{
		PREPAID: ChangeLoadbalancerChargeModeRequestBodyChargeMode{
			value: "prepaid",
		},
	}
}

func (c ChangeLoadbalancerChargeModeRequestBodyChargeMode) Value() string {
	return c.value
}

func (c ChangeLoadbalancerChargeModeRequestBodyChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ChangeLoadbalancerChargeModeRequestBodyChargeMode) UnmarshalJSON(b []byte) error {
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
