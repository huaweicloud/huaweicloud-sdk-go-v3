package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeriodOrderInfo 包年/包月信息体
type PeriodOrderInfo struct {

	// **参数解释**：  订购周期类型。  **约束限制**：  不涉及。  **取值范围**：  - 2：周期类型为月 - 3：周期类型为年  **默认取值**：  不涉及
	PeriodType *PeriodOrderInfoPeriodType `json:"period_type,omitempty"`

	// **参数解释**：  订购周期数。  **约束限制**：  不涉及。  **取值范围**：  - period_type为2（月）时，取值为1-9 - period_type为3（年）时，取值为1-3  **默认取值**：  不涉及
	PeriodNum *int32 `json:"period_num,omitempty"`

	// **参数解释**：  是否自动续订。  **约束限制**：  不涉及。  **取值范围**：  - 0：不自动续订 - 1：自动续订  **默认取值**：  不涉及
	IsAutoRenew *PeriodOrderInfoIsAutoRenew `json:"is_auto_renew,omitempty"`
}

func (o PeriodOrderInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PeriodOrderInfo struct{}"
	}

	return strings.Join([]string{"PeriodOrderInfo", string(data)}, " ")
}

type PeriodOrderInfoPeriodType struct {
	value int32
}

type PeriodOrderInfoPeriodTypeEnum struct {
	E_2 PeriodOrderInfoPeriodType
	E_3 PeriodOrderInfoPeriodType
}

func GetPeriodOrderInfoPeriodTypeEnum() PeriodOrderInfoPeriodTypeEnum {
	return PeriodOrderInfoPeriodTypeEnum{
		E_2: PeriodOrderInfoPeriodType{
			value: 2,
		}, E_3: PeriodOrderInfoPeriodType{
			value: 3,
		},
	}
}

func (c PeriodOrderInfoPeriodType) Value() int32 {
	return c.value
}

func (c PeriodOrderInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodOrderInfoPeriodType) UnmarshalJSON(b []byte) error {
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

type PeriodOrderInfoIsAutoRenew struct {
	value int32
}

type PeriodOrderInfoIsAutoRenewEnum struct {
	E_0 PeriodOrderInfoIsAutoRenew
	E_1 PeriodOrderInfoIsAutoRenew
}

func GetPeriodOrderInfoIsAutoRenewEnum() PeriodOrderInfoIsAutoRenewEnum {
	return PeriodOrderInfoIsAutoRenewEnum{
		E_0: PeriodOrderInfoIsAutoRenew{
			value: 0,
		}, E_1: PeriodOrderInfoIsAutoRenew{
			value: 1,
		},
	}
}

func (c PeriodOrderInfoIsAutoRenew) Value() int32 {
	return c.value
}

func (c PeriodOrderInfoIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PeriodOrderInfoIsAutoRenew) UnmarshalJSON(b []byte) error {
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
