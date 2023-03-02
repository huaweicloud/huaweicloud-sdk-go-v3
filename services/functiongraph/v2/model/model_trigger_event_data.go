package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TriggerEventData struct {

	// 串行处理数据
	IsSerial *bool `json:"is_serial,omitempty"`

	// 最大字节数
	MaxFetchBytes *int32 `json:"max_fetch_bytes,omitempty"`

	// 拉取周期
	PollingInterval *int32 `json:"polling_interval,omitempty"`

	// 拉取周期单位
	PollingUnit *TriggerEventDataPollingUnit `json:"polling_unit,omitempty"`
}

func (o TriggerEventData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerEventData struct{}"
	}

	return strings.Join([]string{"TriggerEventData", string(data)}, " ")
}

type TriggerEventDataPollingUnit struct {
	value string
}

type TriggerEventDataPollingUnitEnum struct {
	MS TriggerEventDataPollingUnit
	S  TriggerEventDataPollingUnit
}

func GetTriggerEventDataPollingUnitEnum() TriggerEventDataPollingUnitEnum {
	return TriggerEventDataPollingUnitEnum{
		MS: TriggerEventDataPollingUnit{
			value: "ms",
		},
		S: TriggerEventDataPollingUnit{
			value: "s",
		},
	}
}

func (c TriggerEventDataPollingUnit) Value() string {
	return c.value
}

func (c TriggerEventDataPollingUnit) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TriggerEventDataPollingUnit) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
