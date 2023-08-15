package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PeriodOrderInfo 包年/包月信息体。
type PeriodOrderInfo struct {

	// 订购周期类型。取值： - 2：月。 - 3：年。
	PeriodType PeriodOrderInfoPeriodType `json:"period_type"`

	// 订购周期数。根据period_type取值不同，代表不同周期数，例如： - 当period_type为2时，period_num为1代表1月。 - 当period_type为3时，period_num为1代表1年。
	PeriodNum int32 `json:"period_num"`

	// 是否自动续订。取值： - 0：否（默认值，需要客户手动去支付）。 - 1：是（自动支付）。
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
