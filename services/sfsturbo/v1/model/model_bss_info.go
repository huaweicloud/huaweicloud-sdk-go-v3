package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 计费参数详情
type BssInfo struct {

	// 是否自动续费
	IsAutoRenew *BssInfoIsAutoRenew `json:"is_auto_renew,omitempty"`

	// 包周期订购的周期数
	PeriodNum int64 `json:"period_num"`

	// 包周期的类型，可选包年或包月，2 表示包月，3 表示包年
	PeriodType BssInfoPeriodType `json:"period_type"`

	// 是否生成订单后自动扣款
	IsAutoPay *BssInfoIsAutoPay `json:"is_auto_pay,omitempty"`
}

func (o BssInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BssInfo struct{}"
	}

	return strings.Join([]string{"BssInfo", string(data)}, " ")
}

type BssInfoIsAutoRenew struct {
	value int64
}

type BssInfoIsAutoRenewEnum struct {
	E_0 BssInfoIsAutoRenew
	E_1 BssInfoIsAutoRenew
}

func GetBssInfoIsAutoRenewEnum() BssInfoIsAutoRenewEnum {
	return BssInfoIsAutoRenewEnum{
		E_0: BssInfoIsAutoRenew{
			value: 0,
		}, E_1: BssInfoIsAutoRenew{
			value: 1,
		},
	}
}

func (c BssInfoIsAutoRenew) Value() int64 {
	return c.value
}

func (c BssInfoIsAutoRenew) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssInfoIsAutoRenew) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}

type BssInfoPeriodType struct {
	value int64
}

type BssInfoPeriodTypeEnum struct {
	E_2 BssInfoPeriodType
	E_3 BssInfoPeriodType
}

func GetBssInfoPeriodTypeEnum() BssInfoPeriodTypeEnum {
	return BssInfoPeriodTypeEnum{
		E_2: BssInfoPeriodType{
			value: 2,
		}, E_3: BssInfoPeriodType{
			value: 3,
		},
	}
}

func (c BssInfoPeriodType) Value() int64 {
	return c.value
}

func (c BssInfoPeriodType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssInfoPeriodType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}

type BssInfoIsAutoPay struct {
	value int64
}

type BssInfoIsAutoPayEnum struct {
	E_0 BssInfoIsAutoPay
	E_1 BssInfoIsAutoPay
}

func GetBssInfoIsAutoPayEnum() BssInfoIsAutoPayEnum {
	return BssInfoIsAutoPayEnum{
		E_0: BssInfoIsAutoPay{
			value: 0,
		}, E_1: BssInfoIsAutoPay{
			value: 1,
		},
	}
}

func (c BssInfoIsAutoPay) Value() int64 {
	return c.value
}

func (c BssInfoIsAutoPay) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BssInfoIsAutoPay) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int64")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int64)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int64 error")
	}
}
