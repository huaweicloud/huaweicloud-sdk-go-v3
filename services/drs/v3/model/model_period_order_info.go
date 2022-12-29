package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 包年/包月信息体
type PeriodOrderInfo struct {

	// 订购周期类型。 取值范围： - 2：表示周期类型为月。 - 3：表示周期类型为年。
	PeriodType *PeriodOrderInfoPeriodType `json:"period_type,omitempty"`

	// 订购周期数。 取值范围： - period_type=2（周期类型为月）时，取值为[1，9]。 - period_type=3（周期类型为年）时，取值为[1，3]。
	PeriodNum *int32 `json:"period_num,omitempty"`

	// 是否自动续订。 取值范围： - 0：表示不自动续订。 - 1：表示自动续订。
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
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
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
