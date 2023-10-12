package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Period 告警条件判断周期，单位为秒，支持的值为1，300，1200，3600，14400，86400。说明：当period设置为1时，表示以原始的指标数据判断告警。当alarm_type为（EVENT.SYS| EVENT.CUSTOM）时允许为0。
type Period struct {
	value int32
}

type PeriodEnum struct {
	E_0     Period
	E_1     Period
	E_300   Period
	E_1200  Period
	E_3600  Period
	E_14400 Period
	E_86400 Period
}

func GetPeriodEnum() PeriodEnum {
	return PeriodEnum{
		E_0: Period{
			value: 0,
		}, E_1: Period{
			value: 1,
		}, E_300: Period{
			value: 300,
		}, E_1200: Period{
			value: 1200,
		}, E_3600: Period{
			value: 3600,
		}, E_14400: Period{
			value: 14400,
		}, E_86400: Period{
			value: 86400,
		},
	}
}

func (c Period) Value() int32 {
	return c.value
}

func (c Period) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Period) UnmarshalJSON(b []byte) error {
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
